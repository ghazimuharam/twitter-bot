package cron

import (
	"fmt"
	"strings"

	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	internal_twitter "github.com/ghazimuharam/twitter-bot/internal/twitter"
	internal_entity "github.com/ghazimuharam/twitter-bot/internal/twitter/entity"
	"github.com/ghazimuharam/twitter-bot/pkg/regex"
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
	if c.configs.App.NumberOfDM == 0 {
		c.configs.App.NumberOfDM = defaultNumberOfDM
	}

	if c.configs.App.TriggerWord == "" {
		c.configs.App.TriggerWord = "Trigger!"
	}

	lastDirectMsgID, err := c.cacheRepo.Get(internal_entity.MessageIDCacheKey)
	if err != nil {
		fmt.Println("cache not found")
	} else {
		fmt.Printf("using %s as lastDirectMsgID\n", lastDirectMsgID)
	}

	dms, err := c.directMsgUC.GetCleanDirectMessages("", lastDirectMsgID, c.configs.App.NumberOfDM)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, dm := range dms.Events {
		// check if we overlapping with tweeted message
		if dm.ID == lastDirectMsgID {
			// if message already tweeted, break the loop
			fmt.Println("dm.ID is overlapping with lastDirectMsgID")
			break
		}

		// check if dm text have a trigger word
		directMsgText := regex.RemoveURLFromString(dm.Message.Data.Text)
		if !strings.Contains(strings.ToLower(directMsgText), strings.ToLower(c.configs.App.TriggerWord)) {
			fmt.Println("dm text doesn't have a trigger word")
			continue
		}

		// make it safe because we need to access the data when
		// we tweet
		mediaUploaded := &twitter.MediaUploadResult{}
		if dm.Message.Data.Attachment != nil {
			mediaByte, err := c.directMsgUC.GetMediaFromDirectMessage(dm.Message.Data.Attachment.Media.MediaURL)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if dm.Message.Data.Attachment.Media.Type == internal_entity.SupportedMediaType {
				mediaUploaded, err = c.mediaUC.Upload(mediaByte, dm.Message.Data.Attachment.Media.Type)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		}

		tweet, err := c.tweetUC.CreateTweet(
			regex.RemoveURLFromString(dm.Message.Data.Text),
			setMediaFile(mediaUploaded),
		)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// logging
		fmt.Println(tweet.IDStr + " - " + tweet.Text)
	}

	// only set cache if dms event available
	if dms != nil && len(dms.Events) != 0 {
		lastDirectMsgID = dms.Events[len(dms.Events)-1].ID
		// set last read message to local cache, only log
		// if lastDirectMsgID is not empty
		if lastDirectMsgID != "" {
			c.cacheRepo.Set(internal_entity.MessageIDCacheKey, lastDirectMsgID)
		}
	}
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
