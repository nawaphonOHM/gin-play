package handler_cache_test

import "github.com/gin-gonic/gin"

type CacheTestHandler interface {
	Get(c *gin.Context)
	Set(c *gin.Context)
}

type SaveCache struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
