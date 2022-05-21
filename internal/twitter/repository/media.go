package repository

import (
	"fmt"

	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	"github.com/ghazimuharam/twitter-bot/internal/twitter/repository/client"
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
	configs *entity.Config
	client  client.TwitterClientWrapperItf
}

// NewMediaRepo to initialize MediaRepo Repository
func NewMediaRepo(configs *entity.Config, client client.TwitterClientWrapperItf) *MediaRepo {
	return &MediaRepo{
		configs: configs,
		client:  client,
	}
}

// Upload to create a new tweet
func (twt *MediaRepo) Upload(media []byte, mediaType string) (*twitter.MediaUploadResult, error) {
	mediaType, ok := mediaTypes[mediaType]
	if !ok {
		return nil, fmt.Errorf("media type not found")
	}

	// upload media using twitter client
	uploadedMedia, _, err := twt.client.UploadMedia(media, mediaType)
	if err != nil {
		return nil, err
	}
	if uploadedMedia == nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	return uploadedMedia, nil
}
