package usecase

import (
	"fmt"

	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	internal_twitter "github.com/ghazimuharam/twitter-bot/internal/twitter"
)

type DirectMessageUsecase struct {
	directMsgRepo internal_twitter.DirectMessageRepoItf
	cacheRepo     internal_twitter.CacheRepoItf
	configs       *entity.Config
}

func NewDirectMessageUsecase(directMsgRepo internal_twitter.DirectMessageRepoItf, cacheRepo internal_twitter.CacheRepoItf, configs *entity.Config) *DirectMessageUsecase {
	return &DirectMessageUsecase{
		directMsgRepo: directMsgRepo,
		cacheRepo:     cacheRepo,
		configs:       configs,
	}
}

func (uc *DirectMessageUsecase) DeleteDirectMessages(directMsgID string) (bool, error) {
	return uc.directMsgRepo.DeleteDirectMessages(directMsgID)
}

func (uc *DirectMessageUsecase) GetDirectMessages(cursor string, numberOfDM int) (*twitter.DirectMessageEvents, error) {
	return uc.directMsgRepo.GetDirectMessages(cursor, numberOfDM)
}

func (uc *DirectMessageUsecase) GetCleanDirectMessages(cursor string, lastDirectMsgID string, numberOfDM int) (*twitter.DirectMessageEvents, error) {
	allDirectMessage := &twitter.DirectMessageEvents{}

	for len(allDirectMessage.Events) < numberOfDM {
		isOverlapDirectMsgID := false
		dms, err := uc.directMsgRepo.GetDirectMessages(cursor, uc.configs.App.DefaultCountTweetRetriever)
		if err != nil {
			return nil, err
		}

		for _, dm := range dms.Events {
			// get only received message
			if dm.Message.SenderID != uc.configs.App.Access.AccountID {
				// check if we overlapping with tweeted message
				if dm.ID == lastDirectMsgID {
					// if message already tweeted, break the loop
					fmt.Println("dm.ID is overlapping with lastDirectMsgID")
					isOverlapDirectMsgID = true
					break
				}

				allDirectMessage.Events = append(allDirectMessage.Events, dm)
			}
		}

		if isOverlapDirectMsgID {
			break
		}

		cursor = dms.NextCursor
	}

	return allDirectMessage, nil
}

func (uc *DirectMessageUsecase) GetMediaFromDirectMessage(mediaURL string) ([]byte, error) {
	return uc.directMsgRepo.GetMediaFromDirectMessage(mediaURL)
}
