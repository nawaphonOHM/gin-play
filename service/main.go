package main

import (
	"gin-play/handlers/cache_test"
	"gin-play/handlers/hello_world"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"strings"
)

func main() {

	redisConfig := &redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}

	if redisPassword, has := os.LookupEnv("REDIS_PASSWORD"); has {
		redisConfig.Password = redisPassword
	}

	if redisHost, has := os.LookupEnv("REDIS_URL"); has {
		redisConfig.Addr = redisHost
	}

	redisClient := redis.NewClient(redisConfig)

	ginDefault := gin.Default()

	helloWorldHandler := hello_world.NewHelloWorldHandler()
	cacheTestHandler := cache_test.NewCacheTestHandler(redisClient)

	root := ginDefault.Group("/")
	{
		root.GET("/hello-world", helloWorldHandler.MainHandler)
	}

	cache := ginDefault.Group("/cache")
	{
		cache.GET("/hello-world/:key", cacheTestHandler.Get)
		cache.POST("/hello-world", cacheTestHandler.Set)
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
