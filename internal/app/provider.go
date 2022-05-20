package app

import (
	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	internal_twitter "github.com/ghazimuharam/twitter-bot/internal/twitter"
	"github.com/ghazimuharam/twitter-bot/internal/twitter/delivery/cron"
	"github.com/ghazimuharam/twitter-bot/internal/twitter/repository"
	"github.com/ghazimuharam/twitter-bot/internal/twitter/repository/client"
	"github.com/ghazimuharam/twitter-bot/internal/twitter/usecase"
	cronuc "github.com/ghazimuharam/twitter-bot/internal/twitter/usecase/cron"
	"github.com/google/wire"
)

var (
	ucSet = wire.NewSet(
		provideDirectMessageUsecase,
		wire.Bind(new(internal_twitter.DirectMessageUCItf), new(*usecase.DirectMessageUsecase)),
		provideTweetUsecase,
		wire.Bind(new(internal_twitter.TweetUCItf), new(*usecase.TweetUsecase)),
		provideMediaUsecase,
		wire.Bind(new(internal_twitter.MediaUCItf), new(*usecase.MediaUsecase)),
		provideCronUsecase,
		wire.Bind(new(internal_twitter.CronUCitf), new(*cronuc.CronUsecase)),
	)

	repoSet = wire.NewSet(
		provideTwitterClient,
		provideDirectMessageRepo,
		wire.Bind(new(internal_twitter.DirectMessageRepoItf), new(*repository.DirectMessageRepo)),
		provideTweetRepo,
		wire.Bind(new(internal_twitter.TweetRepoItf), new(*repository.TweetRepo)),
		provideMediaRepo,
		wire.Bind(new(internal_twitter.MediaRepoItf), new(*repository.MediaRepo)),
		provideCacheRepo,
		wire.Bind(new(internal_twitter.CacheRepoItf), new(*repository.CacheRepo)),
	)

	allSet = wire.NewSet(
		ucSet,
		repoSet,
		provideCronApplication,
	)
)

func provideTwitterClient(configs *entity.Config) *twitter.Client {
	// http.Client will automatically authorize Requests
	httpClient := client.InitOauthClient(configs)

	// Twitter client
	return twitter.NewClient(httpClient)
}

func provideDirectMessageRepo(configs *entity.Config, client *twitter.Client) *repository.DirectMessageRepo {
	return repository.NewDirectMessage(configs, client)
}

func provideTweetRepo(client *twitter.Client, configs *entity.Config) *repository.TweetRepo {
	return repository.NewTweetRepo(client, configs)
}

func provideMediaRepo(client *twitter.Client, configs *entity.Config) *repository.MediaRepo {
	return repository.NewMediaRepo(client, configs)
}

func provideCacheRepo(configs *entity.Config) *repository.CacheRepo {
	return repository.NewCacheRepo(configs)
}

func provideDirectMessageUsecase(
	directMsgrepo internal_twitter.DirectMessageRepoItf,
	cacheRepo internal_twitter.CacheRepoItf,
	configs *entity.Config) *usecase.DirectMessageUsecase {
	return usecase.NewDirectMessageUsecase(directMsgrepo, cacheRepo, configs)
}

func provideTweetUsecase(repo internal_twitter.TweetRepoItf) *usecase.TweetUsecase {
	return usecase.NewTweetUsecase(repo)
}

func provideMediaUsecase(repo internal_twitter.MediaRepoItf) *usecase.MediaUsecase {
	return usecase.NewMediaUsecase(repo)
}

func provideCronUsecase(
	configs *entity.Config,
	directMsgUC internal_twitter.DirectMessageUCItf,
	tweetUC internal_twitter.TweetUCItf,
	mediaUC internal_twitter.MediaUCItf,
	cacheRepo internal_twitter.CacheRepoItf,
) *cronuc.CronUsecase {
	return cronuc.NewCronUsecase(configs, directMsgUC, tweetUC, mediaUC, cacheRepo)
}

func provideCronApplication(uc internal_twitter.CronUCitf, configs *entity.Config) *cron.CronApplication {
	return cron.NewCronApplication(uc, configs)
}
