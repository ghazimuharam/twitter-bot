package twitter

import (
	"github.com/ghazimuharam/go-twitter/twitter"
)

type DirectMessageRepoItf interface {
	DeleteDirectMessages(directMsgID string) (bool, error)
	GetDirectMessages(cursor string, numberOfDM int) (*twitter.DirectMessageEvents, error)
	GetMediaFromDirectMessage(mediaURL string) ([]byte, error)
}

type TweetRepoItf interface {
	CreateTweet(tweet string, mediaIds []int64) (*twitter.Tweet, error)
}

type MediaRepoItf interface {
	Upload(media []byte, mediaType string) (*twitter.MediaUploadResult, error)
}

type CacheRepoItf interface {
	Get(key string) (string, error)
	Set(key string, value interface{})
}
