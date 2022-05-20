package client

import (
	"net/http"

	"github.com/dghubble/oauth1"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
)

func InitOauthClient(configs *entity.Config) *http.Client {
	// OAuth1
	config := oauth1.NewConfig(configs.App.Consumer.Key, configs.App.Consumer.Secret)
	token := oauth1.NewToken(configs.App.Access.Token, configs.App.Access.Secret)
	// http.Client will automatically authorize Requests
	return config.Client(oauth1.NoContext, token)
}
