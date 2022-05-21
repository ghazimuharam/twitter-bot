package twitter

import (
	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/internal/twitter/entity"
)

type DirectMessageRepoItf interface {
	DeleteDirectMessages(directMsgID string) (bool, error)
	GetDirectMessages(cursor string) (*twitter.DirectMessageEvents, error)
	GetMediaFromDirectMessage(mediaURL string) ([]byte, error)
	SendDirectMessage(tweet, recipientID string) (*twitter.DirectMessageEvent, error)
}

type TweetRepoItf interface {
	CreateTweet(entity.Tweet) (*twitter.Tweet, error)
}

type MediaRepoItf interface {
	Upload(media []byte, mediaType string) (*twitter.MediaUploadResult, error)
}

type CacheRepoItf interface {
	Get(key string) (string, error)
	Set(key string, value interface{})
}
