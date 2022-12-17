package cron

import (
	"fmt"
	"log"
	"strings"

	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	internal_twitter "github.com/ghazimuharam/twitter-bot/internal/twitter"
	internal_entity "github.com/ghazimuharam/twitter-bot/internal/twitter/entity"
	"github.com/ghazimuharam/twitter-bot/pkg/helper"
	"github.com/ghazimuharam/twitter-bot/pkg/regex"
	"github.com/ghazimuharam/twitter-bot/pkg/resolver"
)

const (
	defaultNumberOfDM = 10
)

type CronUsecase struct {
	configs     *entity.Config
	directMsgUC internal_twitter.DirectMessageUCItf
	tweetUC     internal_twitter.TweetRepoItf
	mediaUC     internal_twitter.MediaUCItf
	cacheRepo   internal_twitter.CacheRepoItf
}

func NewCronUsecase(
	configs *entity.Config,
	directMsgUC internal_twitter.DirectMessageUCItf,
	tweetUC internal_twitter.TweetUCItf,
	mediaUC internal_twitter.MediaUCItf,
	cacheRepo internal_twitter.CacheRepoItf,
) *CronUsecase {
	return &CronUsecase{
		directMsgUC: directMsgUC,
		configs:     configs,
		tweetUC:     tweetUC,
		cacheRepo:   cacheRepo,
		mediaUC:     mediaUC,
	}
}

func (c *CronUsecase) TweetFromDirectMessage() {
	// define number of dm to retrieve, if the number
	// of dm is 0 from the configs, change it to defaultNumberOfDM
	if c.configs.App.NumberOfDM == 0 {
		c.configs.App.NumberOfDM = defaultNumberOfDM
	}

	// define automatic trigger word, if the trigger word
	// not defined in configs, set it to Trigger!
	if c.configs.App.TriggerWord == "" {
		c.configs.App.TriggerWord = "Trigger!"
	}

	/*
		line deprecated because other solution is provided

		// last direct msg id used to keep up
		// lastDirectMsgID, err := c.cacheRepo.Get(internal_entity.MessageIDCacheKey)
		// if err != nil {
		// 	fmt.Println("cache not found")
		// } else {
		// 	fmt.Printf("using %s as lastDirectMsgID\n", lastDirectMsgID)
		// }
	*/

	// get direct message from user
	dms, err := c.directMsgUC.GetCleanDirectMessages("", "", c.configs.App.NumberOfDM)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, dm := range dms.Events {
		var toTweet internal_entity.Tweet

		/*
			line deprecated because other solution is provided

			// check if we overlapping with tweeted message
			// if dm.ID == lastDirectMsgID {
			// 	// if message already tweeted, break the loop
			// 	fmt.Println("dm.ID is overlapping with lastDirectMsgID")
			// 	break
			// }
		*/

		// check if dm text have a trigger word
		directMsgText := regex.RemoveURLFromString(dm.Message.Data.Text)
		toTweet.Tweet = directMsgText

		if !c.isContainsTriggerWord(directMsgText) || strings.ToLower(directMsgText) == strings.ToLower(c.configs.App.TriggerWord) {
			// delete the dm if it's doesn't contains any trigger word or contains only trigger word
			_, err := c.directMsgUC.DeleteDirectMessages(dm.ID)
			if err != nil {
				fmt.Println(err)
			}

			log.Printf("dm.ID %s is not contains trigger word or contains only trigger word, deleted\n", dm.ID)
			continue
		}

		// check is there any url in dm
		if urlInDM := regex.MatchURLFromString(dm.Message.Data.Text); urlInDM != "" {
			// set the url that have been resolved as a tweet attachment url
			toTweet.AttachmentURL, err = resolver.GetRedirectURL(urlInDM)
			if err != nil {
				fmt.Println("error resolving: ", err)
				continue
			}
		}

		// if tweet attachment url is not related to twitter
		// remove it from attachment and use the text from dm
		if !strings.Contains(toTweet.AttachmentURL, "twitter.com") {
			toTweet.AttachmentURL = ""
			toTweet.Tweet = dm.Message.Data.Text
		}

		// define media uploaded so it will not cause a panic
		mediaUploaded := &twitter.MediaUploadResult{}

		// check dm attachment, if attachment is presence
		// the value is not nil
		if dm.Message.Data.Attachment != nil {
			// get media byte from dm attachment
			mediaByte, err := c.directMsgUC.GetMediaFromDirectMessage(dm.Message.Data.Attachment.Media.MediaURL)
			if err != nil {
				fmt.Println(err)
				continue
			}

			// check is media downloaded is in a list of supported media
			if dm.Message.Data.Attachment.Media.Type == internal_entity.SupportedMediaType {
				// upload media to twitter
				mediaUploaded, err = c.mediaUC.Upload(mediaByte, dm.Message.Data.Attachment.Media.Type)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}

			// remove any attachment url if
			toTweet.AttachmentURL = ""
		}

		// set media ids from uploaded media
		toTweet.MediaIds = setMediaFile(mediaUploaded)

		// create a tweet
		tweet, err := c.tweetUC.CreateTweet(toTweet)
		if err != nil {
			fmt.Println(err)
			continue
		}

		_, err = c.directMsgUC.SendDirectMessage("Success "+helper.TwitterURLBuilder(
			c.configs.App.Account.Handler,
			tweet.IDStr,
		), dm.Message.SenderID)
		if err != nil {
			fmt.Println(err)
		}

		// delete the dm if it's already tweeted
		isDeleted, err := c.directMsgUC.DeleteDirectMessages(dm.ID)
		if err != nil {
			fmt.Println(err)
		}

		// log if all process is clear
		if isDeleted {
			fmt.Println(tweet.IDStr + " - " + tweet.Text)
		}
	}

	/*
		line deprecated because other solution is provided

		// only set cache if dms event available
		// if dms != nil && len(dms.Events) != 0 {
		// 	lastDirectMsgID = dms.Events[len(dms.Events)-1].ID
		// 	// set last read message to local cache, only log
		// 	// if lastDirectMsgID is not empty
		// 	if lastDirectMsgID != "" {
		// 		c.cacheRepo.Set(internal_entity.MessageIDCacheKey, lastDirectMsgID)
		// 	}
		// }
	*/
}

func (c *CronUsecase) isContainsTriggerWord(directMsgText string) bool {
	return strings.Contains(strings.ToLower(directMsgText), strings.ToLower(c.configs.App.TriggerWord))
}

func setMediaFile(mediaUploaded *twitter.MediaUploadResult) []int64 {
	// if no media uploaded, return empty array of int64
	if mediaUploaded.MediaID == 0 {
		return []int64{}
	}

	return []int64{
		mediaUploaded.MediaID,
	}
}
