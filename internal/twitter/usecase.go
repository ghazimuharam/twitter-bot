package twitter

import (
	"github.com/ghazimuharam/go-twitter/twitter"
)

type TweetUCItf interface {
	CreateTweet(tweet string, mediaIds []int64) (*twitter.Tweet, error)
}

type MediaUCItf interface {
	Upload(media []byte, mediaType string) (*twitter.MediaUploadResult, error)
}

type DirectMessageUCItf interface {
	GetDirectMessages(cursor string, numberOfDM int) (*twitter.DirectMessageEvents, error)
	GetCleanDirectMessages(cursor string, lastDirectMsgID string, numberOfDM int) (*twitter.DirectMessageEvents, error)
	GetMediaFromDirectMessage(mediaURL string) ([]byte, error)
}

type CronUCitf interface {
	TweetFromDirectMessage()
}
