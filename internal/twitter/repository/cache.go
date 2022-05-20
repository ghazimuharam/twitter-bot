package repository

import (
	"fmt"

	"github.com/ghazimuharam/twitter-bot/cmd/config/entity"
	"github.com/patrickmn/go-cache"
)

type CacheRepo struct {
	configs *entity.Config
	cache   *cache.Cache
}

func NewCacheRepo(configs *entity.Config) *CacheRepo {
	c := cache.New(configs.Vendor.LocalCache.DefaultExpiration, configs.Vendor.LocalCache.CleanupInterval)

	return &CacheRepo{
		configs: configs,
		cache:   c,
	}
}

func (c *CacheRepo) Get(key string) (string, error) {
	foo, found := c.cache.Get(key)
	if found {
		return foo.(string), nil
	}

	return "", fmt.Errorf("cache not found")
}

func (c *CacheRepo) Set(key string, value interface{}) {
	c.cache.Set(key, value, cache.NoExpiration)
}
