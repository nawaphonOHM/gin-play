package main

import (
	"gin-play/handlers/hello_world"
	"github.com/gin-gonic/gin"
)

func main() {
	ginDefault := gin.Default()

	helloWorldHandler := hello_world.NewHelloWorldHandler()

	ginDefault.Group("/")
	{
		ginDefault.GET("/hello-world", helloWorldHandler.MainHandler)
	}

	err := ginDefault.Run(":8080")
	if err != nil {
		panic("Unable to start server")
	}
}
