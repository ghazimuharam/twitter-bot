package repository

import (
	"fmt"

	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
)

var (
	mediaTypes = map[string]string{
		"photo": "image/jpeg",
		"video": "video/mp4",
	}
	mediaCategories = map[string]string{
		"image/jpeg": "tweet_image",
		"video/mp4":  "tweet_video",
	}
)

type MediaRepo struct {
	client  *twitter.Client
	configs *entity.Config
}

// NewMediaRepo to initialize MediaRepo Repository
func NewMediaRepo(client *twitter.Client, configs *entity.Config) *MediaRepo {
	return &MediaRepo{
		client:  client,
		configs: configs,
	}
}

// Upload to create a new tweet
func (twt *MediaRepo) Upload(media []byte, mediaType string) (*twitter.MediaUploadResult, error) {
	mediaType, ok := mediaTypes[mediaType]
	if !ok {
		return nil, fmt.Errorf("media type not found")
	}

	mediaCategory, ok := mediaCategories[mediaType]
	if !ok {
		return nil, fmt.Errorf("media category not found")
	}

	// upload media using twitter client
	uploadedMedia, _, err := twt.client.Media.Upload(media, mediaType, mediaCategory)
	if err != nil {
		return nil, err
	}
	if uploadedMedia == nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	return uploadedMedia, nil
}
