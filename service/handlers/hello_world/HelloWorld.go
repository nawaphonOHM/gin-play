package hello_world

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Response struct {
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

type HelloWorld interface {
	MainHandler(c *gin.Context)
}
