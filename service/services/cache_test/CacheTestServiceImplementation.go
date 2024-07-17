package cache_test

import "github.com/redis/go-redis/v9"

type CacheTestServiceImplementation struct {
	cache *redis.Client
}

func (s *CacheTestServiceImplementation) Get(key string) string {
	//TODO implement me
	panic("implement me")
}

func NewCacheTestServiceImplementation(cache *redis.Client) CacheTestService {
	return &CacheTestServiceImplementation{cache: cache}
}
