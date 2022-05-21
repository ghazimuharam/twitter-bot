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
	repoSet = wire.NewSet(
		provideTwitterClient,
		provideTwitterClientWrapper,
		wire.Bind(new(client.TwitterClientWrapperItf), new(*client.TwitterClientWrapper)),
		provideDirectMessageRepo,
		wire.Bind(new(internal_twitter.DirectMessageRepoItf), new(*repository.DirectMessageRepo)),
		provideTweetRepo,
		wire.Bind(new(internal_twitter.TweetRepoItf), new(*repository.TweetRepo)),
		provideMediaRepo,
		wire.Bind(new(internal_twitter.MediaRepoItf), new(*repository.MediaRepo)),
		provideCacheRepo,
		wire.Bind(new(internal_twitter.CacheRepoItf), new(*repository.CacheRepo)),
	)

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

func provideDirectMessageRepo(configs *entity.Config, client client.TwitterClientWrapperItf) *repository.DirectMessageRepo {
	return repository.NewDirectMessageRepo(configs, client)
}

func provideTweetRepo(configs *entity.Config, client client.TwitterClientWrapperItf) *repository.TweetRepo {
	return repository.NewTweetRepo(configs, client)
}

func provideMediaRepo(configs *entity.Config, client client.TwitterClientWrapperItf) *repository.MediaRepo {
	return repository.NewMediaRepo(configs, client)
}

func provideTwitterClientWrapper(c *twitter.Client) *client.TwitterClientWrapper {
	return client.NewTwitterClientWrapper(c)
}

func provideCacheRepo(configs *entity.Config) *repository.CacheRepo {
	return repository.NewCacheRepo(configs)
}

func provideDirectMessageUsecase(
	configs *entity.Config,
	directMsgrepo internal_twitter.DirectMessageRepoItf,
	cacheRepo internal_twitter.CacheRepoItf,
) *usecase.DirectMessageUsecase {
	return usecase.NewDirectMessageUsecase(configs, directMsgrepo, cacheRepo)
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
