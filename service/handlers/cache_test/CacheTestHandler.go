package cache_test

import "github.com/gin-gonic/gin"

type CacheTestHandler interface {
	Get(c *gin.Context)
	Set(c *gin.Context)
}
