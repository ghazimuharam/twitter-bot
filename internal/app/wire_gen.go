// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	"github.com/ghazimuharam/twitter-bot/internal/twitter/delivery/cron"
)

// Injectors from wire.go:

func InitCronApp(configs *entity.Config) *cron.CronApplication {
	client := provideTwitterClient(configs)
	twitterClientWrapper := provideTwitterClientWrapper(client)
	directMessageRepo := provideDirectMessageRepo(configs, twitterClientWrapper)
	cacheRepo := provideCacheRepo(configs)
	directMessageUsecase := provideDirectMessageUsecase(configs, directMessageRepo, cacheRepo)
	tweetRepo := provideTweetRepo(configs, twitterClientWrapper)
	tweetUsecase := provideTweetUsecase(tweetRepo)
	mediaRepo := provideMediaRepo(configs, twitterClientWrapper)
	mediaUsecase := provideMediaUsecase(mediaRepo)
	cronUsecase := provideCronUsecase(configs, directMessageUsecase, tweetUsecase, mediaUsecase, cacheRepo)
	cronApplication := provideCronApplication(cronUsecase, configs)
	return cronApplication
}
