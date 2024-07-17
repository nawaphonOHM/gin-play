package handler_database_test

import (
	"github.com/gin-gonic/gin"
)

type DatabaseTestHandlerImplementation struct {
}

func (h *DatabaseTestHandlerImplementation) Get(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewDatabaseTestHandlerImplementation() DatabaseTestHandler {
	return &DatabaseTestHandlerImplementation{}
}
