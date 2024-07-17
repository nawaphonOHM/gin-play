package handler_cache_test

import (
	"gin-play/services/cache_test"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CacheTestHandlerImplementation struct {
	service service_cache_test.CacheTestService
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

// Set @Summary Set value in cache
// @Description Set value in cache with key
// @Tags CacheTestHandler
// @Accept  json
// @Produce  json
// @Param request body handler_cache_test.SaveCache true "KeyValue"
// @Success 200 {object} map[string]interface{} "message":"Cache set successfully"
// @Router /cache/hello-world  [post]
func (h *CacheTestHandlerImplementation) Set(c *gin.Context) {
	body := &SaveCache{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.service.Set(body.Key, body.Value)
	c.JSON(http.StatusOK, gin.H{"message": "Cache set successfully"})
}

func NewCacheTestHandler(service service_cache_test.CacheTestService) CacheTestHandler {
	return &CacheTestHandlerImplementation{service: service}
}
