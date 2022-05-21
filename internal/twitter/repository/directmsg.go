package repository

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	"github.com/ghazimuharam/twitter-bot/internal/twitter/repository/client"
)

// DirectMessageRepo is a struct to hold Direct Message implementation repository
type DirectMessageRepo struct {
	configs *entity.Config
	client  client.TwitterClientWrapperItf
}

// NewDirectMessage to initialize DirectMessage Repository
func NewDirectMessageRepo(configs *entity.Config, client client.TwitterClientWrapperItf) *DirectMessageRepo {
	return &DirectMessageRepo{
		configs: configs,
		client:  client,
	}
}

// DeleteDirectMessages to destroy direct message events
func (dm *DirectMessageRepo) DeleteDirectMessages(directMsgID string) (bool, error) {
	// get direct message using twitter client
	destroyedDirectMsg, err := dm.client.DeleteDirectMessages(directMsgID)
	if err != nil {
		return false, err
	}
	if destroyedDirectMsg == nil {
		return false, fmt.Errorf("err: %v", err)
	}

	return true, nil
}

// GetDirectMessage to get list of direct message including received and sent
func (dm *DirectMessageRepo) GetDirectMessages(cursor string) (*twitter.DirectMessageEvents, error) {
	// get direct message using twitter client
	directMessages, _, err := dm.client.GetDirectMessages(&twitter.DirectMessageEventsListParams{
		Cursor: cursor,
		Count:  dm.configs.App.DefaultCountTweetRetriever,
	})
	if err != nil {
		return nil, err
	}
	if directMessages == nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	return directMessages, nil
}

// GetDirectMessage to get list of direct message including received and sent
func (dm *DirectMessageRepo) GetMediaFromDirectMessage(mediaURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", mediaURL, nil)
	if err != nil {
		return nil, err
	}

	client := client.InitOauthClient(dm.configs)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// SendDirectMessages to sent a direct message
func (dm *DirectMessageRepo) SendDirectMessage(tweet, recipientID string) (*twitter.DirectMessageEvent, error) {
	// get direct message using twitter client
	SendDirectMessages, _, err := dm.client.SendDirectMessage(&twitter.DirectMessageEventsNewParams{
		Event: &twitter.DirectMessageEvent{
			Type: "message_create",
			Message: &twitter.DirectMessageEventMessage{
				Target: &twitter.DirectMessageTarget{
					RecipientID: recipientID,
				},
				Data: &twitter.DirectMessageData{
					Text: tweet,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if SendDirectMessages == nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	return SendDirectMessages, nil
}
