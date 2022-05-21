package twitter

import (
	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/internal/twitter/entity"
)

type TweetUCItf interface {
	CreateTweet(entity.Tweet) (*twitter.Tweet, error)
}

type MediaUCItf interface {
	Upload(media []byte, mediaType string) (*twitter.MediaUploadResult, error)
}

type DirectMessageUCItf interface {
	DeleteDirectMessages(directMsgID string) (bool, error)
	GetCleanDirectMessages(cursor string, lastDirectMsgID string, numberOfDM int) (*twitter.DirectMessageEvents, error)
	GetDirectMessages(cursor string) (*twitter.DirectMessageEvents, error)
	GetMediaFromDirectMessage(mediaURL string) ([]byte, error)
	SendDirectMessage(tweet, recipientID string) (*twitter.DirectMessageEvent, error)
}

type CronUCitf interface {
	TweetFromDirectMessage()
}
