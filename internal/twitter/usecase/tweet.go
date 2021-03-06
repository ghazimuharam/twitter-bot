package usecase

import (
	"github.com/ghazimuharam/go-twitter/twitter"
	internal_twitter "github.com/ghazimuharam/twitter-bot/internal/twitter"
	"github.com/ghazimuharam/twitter-bot/internal/twitter/entity"
)

type TweetUsecase struct {
	repo internal_twitter.TweetRepoItf
}

func NewTweetUsecase(repo internal_twitter.TweetRepoItf) *TweetUsecase {
	return &TweetUsecase{
		repo: repo,
	}
}

// CreateTweet to create a new tweet
func (twt *TweetUsecase) CreateTweet(tweet entity.Tweet) (*twitter.Tweet, error) {
	// get direct message using twitter client
	postedTweet, err := twt.repo.CreateTweet(tweet)
	if err != nil {
		return nil, err
	}

	return postedTweet, nil
}
