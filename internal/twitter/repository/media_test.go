package repository

import (
	"errors"
	"net/http"
	"testing"

	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	"github.com/ghazimuharam/twitter-bot/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewMediaRepo(t *testing.T) {
	t.Run("init media repo", func(t *testing.T) {
		got := NewMediaRepo(nil, nil)
		assert.NotNil(t, got)
	})
}

func TestMediaRepo_Upload(t *testing.T) {
	type mClient struct {
		result *twitter.MediaUploadResult
		resp   *http.Response
		err    error
	}
	type args struct {
		media     []byte
		mediaType string
	}
	tests := []struct {
		name    string
		args    args
		mClient mClient
		want    *twitter.MediaUploadResult
		wantErr bool
	}{
		{
			name: "success upload media",
			args: args{
				media:     []byte{},
				mediaType: "photo",
			},
			mClient: mClient{
				result: &twitter.MediaUploadResult{},
			},
			want: &twitter.MediaUploadResult{},
		},
		{
			name: "failed upload media",
			args: args{
				media:     []byte{},
				mediaType: "video",
			},
			mClient: mClient{
				err: errors.New("something"),
			},
			wantErr: true,
		},
		{
			name: "nil struct from client, return err",
			args: args{
				media:     []byte{},
				mediaType: "video",
			},
			mClient: mClient{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mClient := mock.NewMockTwitterClientWrapperItf(ctrl)
			mClient.EXPECT().
				UploadMedia(gomock.Any(), gomock.Any()).
				Return(tt.mClient.result, tt.mClient.resp, tt.mClient.err)

			mRepo := MediaRepo{
				client:  mClient,
				configs: &entity.Config{},
			}

			got, err := mRepo.Upload(tt.args.media, tt.args.mediaType)
			if (err != nil) != tt.wantErr {
				t.Errorf("MediaRepo.Upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
