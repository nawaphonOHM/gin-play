package handler_database_test

import (
	service_database_test "gin-play/services/database_test"
	"github.com/gin-gonic/gin"
)

type DatabaseTestHandlerImplementation struct {
	service service_database_test.DatabaseTestService
}

func (h *DatabaseTestHandlerImplementation) Get(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewDatabaseTestHandlerImplementation(service service_database_test.DatabaseTestService) DatabaseTestHandler {
	return &DatabaseTestHandlerImplementation{service: service}
}
