package cron

import (
	"reflect"
	"testing"

	"github.com/ghazimuharam/go-twitter/twitter"
	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
)

func Test_setMediaFile(t *testing.T) {
	type args struct {
		mediaUploaded *twitter.MediaUploadResult
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "no media uploaded, return empty array of int64",
			args: args{
				mediaUploaded: &twitter.MediaUploadResult{},
			},
			want: []int64{},
		},
		{
			name: "media uploaded, return array of int64",
			args: args{
				mediaUploaded: &twitter.MediaUploadResult{
					MediaID: 1,
				},
			},
			want: []int64{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setMediaFile(tt.args.mediaUploaded); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setMediaFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCronUsecase_TweetFromDirectMessage(t *testing.T) {
	t.Run("cron running", func(t *testing.T) {
		// tt.c.TweetFromDirectMessage()
	})
}

func TestCronUsecase_isContainsTriggerWord(t *testing.T) {
	type args struct {
		directMsgText string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "direct message text doesn't contains trigger word",
			args: args{
				directMsgText: "test",
			},
			want: false,
		},
		{
			name: "direct message text contains trigger word",
			args: args{
				directMsgText: "!val halo",
			},
			want: true,
		},
		{
			name: "direct message text is trigger word only",
			args: args{
				directMsgText: "!val",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CronUsecase{
				configs: &entity.Config{
					App: entity.App{
						TriggerWord: "!val",
					},
				},
			}
			if got := c.isContainsTriggerWord(tt.args.directMsgText); got != tt.want {
				t.Errorf("CronUsecase.isContainsTriggerWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
