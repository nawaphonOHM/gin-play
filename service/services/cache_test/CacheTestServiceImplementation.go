package cache_test

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

type CacheTestServiceImplementation struct {
	cache *redis.Client
}

func (s *CacheTestServiceImplementation) Get(key string) string {
	log.Println("Get key:", key)

	result, err := s.cache.Get(context.Background(), key).Result()
	if err != nil {
		log.Println("Get err:", err)
		result = ""
	}

	return result
}

func NewCacheTestServiceImplementation(cache *redis.Client) CacheTestService {
	return &CacheTestServiceImplementation{cache: cache}
}
