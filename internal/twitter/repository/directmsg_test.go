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

func TestNewDirectMessage(t *testing.T) {
	t.Run("init", func(t *testing.T) {
		got := NewDirectMessageRepo(nil, nil)
		assert.NotNil(t, got)
	})
}

func TestDirectMessageRepo_DeleteDirectMessages(t *testing.T) {
	type mClient struct {
		resp *http.Response
		err  error
	}
	type args struct {
		directMsgID string
	}
	tests := []struct {
		name    string
		args    args
		mClient mClient
		want    bool
		wantErr bool
	}{
		{
			name: "success delete event, return true",
			args: args{
				directMsgID: "1063573894173323269",
			},
			mClient: mClient{
				resp: &http.Response{},
			},
			want: true,
		},
		{
			name: "nil return http response",
			args: args{
				directMsgID: "1063573894173323269",
			},
			wantErr: true,
		},
		{
			name: "failed delete event, return false",
			args: args{
				directMsgID: "1063573894173323269",
			},
			mClient: mClient{
				resp: &http.Response{},
				err:  errors.New("cannot delete event"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mClient := mock.NewMockTwitterClientWrapperItf(ctrl)
			mClient.EXPECT().
				DeleteDirectMessages(gomock.Any()).
				Return(tt.mClient.resp, tt.mClient.err)

			mRepo := DirectMessageRepo{
				client: mClient,
			}

			got, err := mRepo.DeleteDirectMessages(tt.args.directMsgID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DirectMessageRepo.DeleteDirectMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DirectMessageRepo.DeleteDirectMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirectMessageRepo_GetDirectMessages(t *testing.T) {
	type mClient struct {
		events *twitter.DirectMessageEvents
		resp   *http.Response
		err    error
	}
	type args struct {
		cursor string
	}
	tests := []struct {
		name    string
		args    args
		mClient mClient
		want    *twitter.DirectMessageEvents
		wantErr bool
	}{
		{
			name: "success get direct message",
			mClient: mClient{
				events: &twitter.DirectMessageEvents{
					Events:     []twitter.DirectMessageEvent{},
					NextCursor: "aKsDphtv8uEr",
				},
			},
			want: &twitter.DirectMessageEvents{
				Events:     []twitter.DirectMessageEvent{},
				NextCursor: "aKsDphtv8uEr",
			},
		},
		{
			name: "failed get direct message",
			mClient: mClient{
				events: &twitter.DirectMessageEvents{
					Events:     []twitter.DirectMessageEvent{},
					NextCursor: "aKsDphtv8uEr",
				},
				err: errors.New("something went wrong"),
			},
			wantErr: true,
		},
		{
			name: "nil struct from client, return err",
			mClient: mClient{
				events: nil,
				resp:   &http.Response{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mClient := mock.NewMockTwitterClientWrapperItf(ctrl)
			mClient.EXPECT().
				GetDirectMessages(gomock.Any()).
				Return(tt.mClient.events, tt.mClient.resp, tt.mClient.err)

			mRepo := DirectMessageRepo{
				client:  mClient,
				configs: &entity.Config{},
			}

			got, err := mRepo.GetDirectMessages(tt.args.cursor)
			if (err != nil) != tt.wantErr {
				t.Errorf("DirectMessageRepo.GetDirectMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDirectMessageRepo_SendDirectMessage(t *testing.T) {
	type mClient struct {
		events *twitter.DirectMessageEvent
		resp   *http.Response
		err    error
	}
	type args struct {
		tweet       string
		recipientID string
	}
	tests := []struct {
		name    string
		args    args
		mClient mClient
		want    *twitter.DirectMessageEvent
		wantErr bool
	}{
		{
			name: "success send direct message",
			args: args{
				tweet:       "test tweet",
				recipientID: "1435148913967656960",
			},
			mClient: mClient{
				events: &twitter.DirectMessageEvent{},
			},
			want: &twitter.DirectMessageEvent{},
		},
		{
			name: "failed send direct message",
			mClient: mClient{
				events: &twitter.DirectMessageEvent{},
				err:    errors.New("something went wrong"),
			},
			wantErr: true,
		},
		{
			name:    "nil struct from client, return err",
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
				SendDirectMessage(gomock.Any()).
				Return(tt.mClient.events, tt.mClient.resp, tt.mClient.err)

			mRepo := DirectMessageRepo{
				client:  mClient,
				configs: &entity.Config{},
			}

			got, err := mRepo.SendDirectMessage(tt.args.tweet, tt.args.recipientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DirectMessageRepo.SendDirectMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
