package client

import (
	"net/http"

	"github.com/dghubble/oauth1"
	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
)

type TwitterClientWrapper struct {
	*twitter.Client
}

type TwitterClientWrapperItf interface {
	CreateTweet(status string, params *twitter.StatusUpdateParams) (*twitter.Tweet, *http.Response, error)
	DeleteDirectMessages(directMsgID string) (*http.Response, error)
	GetDirectMessages(params *twitter.DirectMessageEventsListParams) (*twitter.DirectMessageEvents, *http.Response, error)
	SendDirectMessage(params *twitter.DirectMessageEventsNewParams) (*twitter.DirectMessageEvent, *http.Response, error)
	UploadMedia(media []byte, mediaType string) (*twitter.MediaUploadResult, *http.Response, error)
}

func NewTwitterClientWrapper(c *twitter.Client) *TwitterClientWrapper {
	return &TwitterClientWrapper{
		c,
	}
}

func (tc *TwitterClientWrapper) CreateTweet(status string, params *twitter.StatusUpdateParams) (*twitter.Tweet, *http.Response, error) {
	return tc.Statuses.Update(status, params)
}

func (tc *TwitterClientWrapper) DeleteDirectMessages(directMsgID string) (*http.Response, error) {
	return tc.DirectMessages.EventsDestroy(directMsgID)
}

func (tc *TwitterClientWrapper) GetDirectMessages(params *twitter.DirectMessageEventsListParams) (*twitter.DirectMessageEvents, *http.Response, error) {
	return tc.DirectMessages.EventsList(params)
}

func (tc *TwitterClientWrapper) SendDirectMessage(params *twitter.DirectMessageEventsNewParams) (*twitter.DirectMessageEvent, *http.Response, error) {
	return tc.DirectMessages.EventsNew(params)
}

func (tc *TwitterClientWrapper) UploadMedia(media []byte, mediaType string) (*twitter.MediaUploadResult, *http.Response, error) {
	return tc.Media.Upload(media, mediaType)
}

func InitOauthClient(configs *entity.Config) *http.Client {
	// OAuth1
	config := oauth1.NewConfig(configs.App.Consumer.Key, configs.App.Consumer.Secret)
	token := oauth1.NewToken(configs.App.Access.Token, configs.App.Access.Secret)
	// http.Client will automatically authorize Requests
	return config.Client(oauth1.NoContext, token)
}
