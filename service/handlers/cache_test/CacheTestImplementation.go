package cache_test

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

type CacheTestHandlerImplementation struct {
	cache *redis.Client
}

func (h *CacheTestHandlerImplementation) Get(c *gin.Context) {
	key := c.Param("key")
	log.Println("Get key:", key)

	result, err := h.cache.Get(context.Background(), key).Result()
	if err != nil {
		result = ""
		log.Println("Get err:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"value": result,
	})
}

func (h *CacheTestHandlerImplementation) Set(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewCacheTestHandler(cache *redis.Client) CacheTestHandler {
	return &CacheTestHandlerImplementation{cache: cache}
}
