package cache_test

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type CacheTestHandlerImplementation struct {
	cache *redis.Client
}

func (h *CacheTestHandlerImplementation) Get(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *CacheTestHandlerImplementation) Set(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewCacheTestHandler(cache *redis.Client) CacheTestHandler {
	return &CacheTestHandlerImplementation{cache: cache}
}
