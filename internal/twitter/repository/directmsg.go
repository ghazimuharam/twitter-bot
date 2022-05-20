package repository

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	"github.com/ghazimuharam/twitter-bot/internal/twitter/repository/client"
)

type DirectMessageRepo struct {
	configs *entity.Config
	client  *twitter.Client
}

// NewDirectMessage to initialize DirectMessage Repository
func NewDirectMessage(configs *entity.Config, client *twitter.Client) *DirectMessageRepo {
	return &DirectMessageRepo{
		configs: configs,
		client:  client,
	}
}

// GetDirectMessage to get list of direct message including received and sent
func (dm *DirectMessageRepo) GetDirectMessages(cursor string, numberOfDM int) (*twitter.DirectMessageEvents, error) {
	// get direct message using twitter client
	directMessages, _, err := dm.client.DirectMessages.EventsList(&twitter.DirectMessageEventsListParams{
		Cursor: cursor,
		Count:  numberOfDM,
	})
	if err != nil {
		return nil, err
	}
	if directMessages == nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	return directMessages, nil
}

// DeleteDirectMessages to destroy direct message events
func (dm *DirectMessageRepo) DeleteDirectMessages(directMsgID string) (bool, error) {
	// get direct message using twitter client
	destroyedDirectMsg, err := dm.client.DirectMessages.EventsDestroy(directMsgID)
	if err != nil {
		return false, err
	}
	if destroyedDirectMsg == nil {
		return false, fmt.Errorf("err: %v", err)
	}

	return true, nil
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
