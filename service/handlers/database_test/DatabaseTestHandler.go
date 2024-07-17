package handler_database_test

import "github.com/gin-gonic/gin"

type DatabaseTestHandler interface {
	Get(c *gin.Context)
}
