package main

import (
	"gin-play/handlers/cache_test"
	"gin-play/handlers/hello_world"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
)

func main() {
	ginDefault := gin.Default()

	helloWorldHandler := hello_world.NewHelloWorldHandler()

	cacheTestHandler := cache_test.NewCacheTestImplementation()

	ginDefault.Group("/")
	{
		ginDefault.GET("/hello-world", helloWorldHandler.MainHandler)
	}

	if port, isSet := os.LookupEnv("PORT"); isSet {
		err := ginDefault.Run(strings.Join([]string{":", port}, ""))
		if err != nil {
			panic("Unable to start server")
		}

	} else {
		log.Println("Port hasn't been set; Use default PORT = 8080")
		err := ginDefault.Run(strings.Join([]string{":", "8080"}, ""))
		if err != nil {
			panic("Unable to start server")
		}
	}
}
