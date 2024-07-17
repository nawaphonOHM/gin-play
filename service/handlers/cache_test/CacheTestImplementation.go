package handler_cache_test

import (
	"gin-play/services/cache_test"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CacheTestHandlerImplementation struct {
	service cache_test.CacheTestService
}

// Get @Summary Get value from cache
// @Description Get value from cache by key
// @Tags CacheTestHandler
// @Accept  json
// @Produce  json
// @Param   key     path    string     true        "Key"
// @Success 200 {object} map[string]interface{} "value":"string"
// @Router /cache/hello-world/{key} [get]
func (h *CacheTestHandlerImplementation) Get(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"value": h.service.Get(c.Param("key")),
	})
}

func (h *CacheTestHandlerImplementation) Set(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewCacheTestHandler(service cache_test.CacheTestService) CacheTestHandler {
	return &CacheTestHandlerImplementation{service: service}
}
