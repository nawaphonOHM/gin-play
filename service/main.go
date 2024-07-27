package main

import (
	"fmt"
	handler_cache_test "gin-play/handlers/cache_test"
	handler_database_test "gin-play/handlers/database_test"
	handler_downloadfile_test "gin-play/handlers/downloadfile_test"
	handler_hello_world "gin-play/handlers/hello_world"
	repository_database_test "gin-play/repositories/database_test"
	service_cache_test "gin-play/services/cache_test"
	service_database_test "gin-play/services/database_test"
	service_downloadfile_test "gin-play/services/downloadfile_test"
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

	dbTestRepository := repository_database_test.NewDatabaseTestRepositoryImplementation(db)

	cacheTestService := service_cache_test.NewCacheTestServiceImplementation(redisClient)
	dbTestService := service_database_test.NewDatabaseTestServiceImplementation(dbTestRepository)

	helloWorldHandler := handler_hello_world.NewHelloWorldHandler()
	cacheTestHandler := handler_cache_test.NewCacheTestHandler(cacheTestService)
	dbTestHandler := handler_database_test.NewDatabaseTestHandlerImplementation(dbTestService)

	downloadFileService := service_downloadfile_test.NewDownloadFileTestServiceImplementation()
	downloadFileHandler := handler_downloadfile_test.NewDownloadFileTestHandlerImplementation(downloadFileService)

	root := ginDefault.Group("/")
	{
		root.GET("/hello-world", helloWorldHandler.MainHandler)
	}

	cache := ginDefault.Group("/cache")
	{
		cache.GET("/hello-world/:key", cacheTestHandler.Get)
		cache.POST("/hello-world", cacheTestHandler.Set)
	}

	dbPath := ginDefault.Group("/db")
	{
		dbPath.GET("/hello-world", dbTestHandler.Get)
	}

	downloadFilePath := ginDefault.Group("/download")
	{
		downloadFilePath.GET("/hello-world", downloadFileHandler.Get)
	}

	if port, isSet := os.LookupEnv("PORT"); isSet {
		errGin := ginDefault.Run(strings.Join([]string{":", port}, ""))
		if errGin != nil {
			panic("Unable to start server")
		}

	} else {
		log.Println("Port hasn't been set; Use default PORT = 8080")
		errGin := ginDefault.Run(strings.Join([]string{":", "8080"}, ""))
		if errGin != nil {
			panic("Unable to start server")
		}
	}
}
