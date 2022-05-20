package usecase

import (
	"github.com/ghazimuharam/go-twitter/twitter"
	internal_twitter "github.com/ghazimuharam/twitter-bot/internal/twitter"
)

type MediaUsecase struct {
	repo internal_twitter.MediaRepoItf
}

func NewMediaUsecase(repo internal_twitter.MediaRepoItf) *MediaUsecase {
	return &MediaUsecase{
		repo: repo,
	}
}

// Upload to create a new tweet
func (twt *MediaUsecase) Upload(media []byte, mediaType string) (*twitter.MediaUploadResult, error) {
	// get direct message using twitter client
	uploadedMedia, err := twt.repo.Upload(media, mediaType)
	if err != nil {
		return nil, err
	}

	return uploadedMedia, nil
}
