package service_cache_test

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"time"
)

type CacheTestServiceImplementation struct {
	cache *redis.Client
}

func (s *CacheTestServiceImplementation) Set(key string, value string) {
	timeStr, _ := os.LookupEnv("REDIS_CACHE_EXPIRE")

	timeDuration, err := time.ParseDuration(timeStr)
	if err != nil {
		log.Println("Unable to parse REDIS_CACHE_EXPIRE set no expire time")
		timeDuration = 0
	}

	s.cache.Set(context.Background(), key, value, timeDuration)
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
