package cache_test

import "github.com/redis/go-redis/v9"

type CacheTestImplementation struct {
	cache *redis.Client
}

func (c *CacheTestImplementation) Get(key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CacheTestImplementation) Set(key string, value string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewCacheTestImplementation(cache *redis.Client) CacheTest {
	return &CacheTestImplementation{cache: cache}
}
