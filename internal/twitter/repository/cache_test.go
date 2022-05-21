package repository

import (
	"testing"
	"time"

	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
)

func TestNewCacheRepo(t *testing.T) {
	t.Run("init cache repo", func(t *testing.T) {
		got := NewCacheRepo(&entity.Config{})
		assert.NotNil(t, got)
	})
}

func TestCacheRepo_Get(t *testing.T) {
	CacheRepoImp := &CacheRepo{
		configs: &entity.Config{},
		cache:   cache.New(time.Duration(1*time.Minute), time.Duration(1*time.Minute)),
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		c       *CacheRepo
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "cache not found, return err",
			c:    CacheRepoImp,
			args: args{
				key: "testCacheKey",
			},
			wantErr: true,
		},
		{
			name: "cache found, return value",
			c:    CacheRepoImp,
			args: args{
				key: "testCacheKey",
			},
			want: "test value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.wantErr {
				tt.c.Set(tt.args.key, tt.want)
			}

			got, err := tt.c.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("CacheRepo.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CacheRepo.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
