package repository

import (
	"fmt"

	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
)

type TweetRepo struct {
	client  *twitter.Client
	configs *entity.Config
}

// NewTweetRepo to initialize TweetRepo Repository
func NewTweetRepo(client *twitter.Client, configs *entity.Config) *TweetRepo {
	return &TweetRepo{
		client:  client,
		configs: configs,
	}
}

// CreateTweet to create a new tweet
func (twt *TweetRepo) CreateTweet(tweet string, mediaIds []int64) (*twitter.Tweet, error) {
	// post a tweet using twitter client
	postedTweet, _, err := twt.client.Statuses.Update(tweet, &twitter.StatusUpdateParams{
		Status:   tweet,
		MediaIds: mediaIds,
	})
	if err != nil {
		return nil, err
	}
	if postedTweet == nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	return postedTweet, nil
}
