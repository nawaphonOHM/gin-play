package handler_downloadfile_test

import "github.com/gin-gonic/gin"

type DownloadFileTestHandler interface {
	Get(c *gin.Context)
}
