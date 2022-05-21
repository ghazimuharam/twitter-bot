package usecase

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	internal_twitter "github.com/ghazimuharam/twitter-bot/internal/twitter"
)

type DirectMessageUsecase struct {
	configs       *entity.Config
	directMsgRepo internal_twitter.DirectMessageRepoItf
	cacheRepo     internal_twitter.CacheRepoItf
}

func NewDirectMessageUsecase(configs *entity.Config, directMsgRepo internal_twitter.DirectMessageRepoItf, cacheRepo internal_twitter.CacheRepoItf) *DirectMessageUsecase {
	return &DirectMessageUsecase{
		configs:       configs,
		directMsgRepo: directMsgRepo,
		cacheRepo:     cacheRepo,
	}
}

func (uc *DirectMessageUsecase) DeleteDirectMessages(directMsgID string) (bool, error) {
	return uc.directMsgRepo.DeleteDirectMessages(directMsgID)
}

func (uc *DirectMessageUsecase) GetCleanDirectMessages(cursor string, lastDirectMsgID string, numberOfDM int) (*twitter.DirectMessageEvents, error) {
	circuitBreaker := false
	allDirectMessage := &twitter.DirectMessageEvents{}

	for len(allDirectMessage.Events) < numberOfDM {
		/*
			line deprecated because other solution is provided

			// isOverlapDirectMsgID := false
		*/
		maxProcessTime := time.Now().Add(time.Duration(-30) * time.Minute).UnixMilli()
		dms, err := uc.directMsgRepo.GetDirectMessages(cursor)
		if err != nil {
			return nil, err
		}

		for _, dm := range dms.Events {

			createdAtTS, err := strconv.ParseInt(dm.CreatedAt, 10, 64)
			if err != nil {
				fmt.Println("error converting timestamp", err)
				continue
			}

			if maxProcessTime > createdAtTS {
				circuitBreaker = true
				break
			}

			// get only received message
			if dm.Message.SenderID != uc.configs.App.Account.ID {
				/*
					line deprecated because other solution is provided

					// check if we overlapping with tweeted message
					// if dm.ID == lastDirectMsgID {
					// 	// if message already tweeted, break the loop
					// 	fmt.Println("dm.ID is overlapping with lastDirectMsgID")
					// 	isOverlapDirectMsgID = true
					// 	break
					// }
				*/

				allDirectMessage.Events = append(allDirectMessage.Events, dm)
			}
		}

		/*
			line deprecated because other solution is provided

			// if isOverlapDirectMsgID {
			// 	break
			// }
		*/

		if circuitBreaker {
			break
		}

		cursor = dms.NextCursor
	}

	return allDirectMessage, nil
}

func (uc *DirectMessageUsecase) GetDirectMessages(cursor string) (*twitter.DirectMessageEvents, error) {
	return uc.directMsgRepo.GetDirectMessages(cursor)
}

func (uc *DirectMessageUsecase) GetMediaFromDirectMessage(mediaURL string) ([]byte, error) {
	return uc.directMsgRepo.GetMediaFromDirectMessage(mediaURL)
}

func (uc *DirectMessageUsecase) SendDirectMessage(tweet, recipientID string) (*twitter.DirectMessageEvent, error) {
	return uc.directMsgRepo.SendDirectMessage(tweet, recipientID)
}
