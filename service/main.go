package main

import (
	"fmt"
	handler_cache_test "gin-play/handlers/cache_test"
	handler_database_test "gin-play/handlers/database_test"
	handler_hello_world "gin-play/handlers/hello_world"
	service_cache_test "gin-play/services/cache_test"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"strings"
)

func main() {

	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "postgres", "test", "localhost", "5432", "public"))
	if err != nil {
		log.Fatalln(err)
	}

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

	cacheTestService := service_cache_test.NewCacheTestServiceImplementation(redisClient)

	helloWorldHandler := handler_hello_world.NewHelloWorldHandler()
	cacheTestHandler := handler_cache_test.NewCacheTestHandler(cacheTestService)

	dbTestHandler := handler_database_test.NewDatabaseTestHandlerImplementation()

	root := ginDefault.Group("/")
	{
		root.GET("/hello-world", helloWorldHandler.MainHandler)
	}

	cache := ginDefault.Group("/cache")
	{
		cache.GET("/hello-world/:key", cacheTestHandler.Get)
		cache.POST("/hello-world", cacheTestHandler.Set)
	}

	db := ginDefault.Group("/db")
	{
		db.GET("/hello-world", dbTestHandler.Get)
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
