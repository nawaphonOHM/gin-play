package cache_test

import "github.com/redis/go-redis/v9"

type CacheTestHandlerImplementation struct {
	cache *redis.Client
}

func (c *CacheTestHandlerImplementation) Get(key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CacheTestHandlerImplementation) Set(key string, value string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewCacheTestHandler(cache *redis.Client) CacheTestHandler {
	return &CacheTestHandlerImplementation{cache: cache}
}
