package hello_world

import (
	"github.com/gin-gonic/gin"
	"time"
)

type HelloWorldHandler struct{}

// MainHandler @Summary Main handlers
// @Description Get Hello World Message
// @ID main-handlers
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "message":"Hello World", "date" : "time.Time"
// @Router /hello-world [get]
// @BasePath /
func (h *HelloWorldHandler) MainHandler(c *gin.Context) {
	response := Response{}

	response.Message = "Hello World"
	response.Date = time.Now()

	c.JSON(200, response)
}

func NewHelloWorldHandler() HelloWorld {
	return &HelloWorldHandler{}
}
