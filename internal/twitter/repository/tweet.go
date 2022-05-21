package repository

import (
	"fmt"

	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	internal_twitter "github.com/ghazimuharam/twitter-bot/internal/twitter/entity"
	"github.com/ghazimuharam/twitter-bot/internal/twitter/repository/client"
)

type TweetRepo struct {
	configs *entity.Config
	client  client.TwitterClientWrapperItf
}

// NewTweetRepo to initialize TweetRepo Repository
func NewTweetRepo(configs *entity.Config, client client.TwitterClientWrapperItf) *TweetRepo {
	return &TweetRepo{
		configs: configs,
		client:  client,
	}
}

// CreateTweet to create a new tweet
func (twt *TweetRepo) CreateTweet(tweet internal_twitter.Tweet) (*twitter.Tweet, error) {
	// post a tweet using twitter client
	postedTweet, _, err := twt.client.CreateTweet(tweet.Tweet, &twitter.StatusUpdateParams{
		Status:        tweet.Tweet,
		MediaIds:      tweet.MediaIds,
		AttachmentURL: tweet.AttachmentURL,
	})
	if err != nil {
		return nil, err
	}
	if postedTweet == nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	return postedTweet, nil
}
