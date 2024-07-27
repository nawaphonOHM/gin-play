package handler_downloadfile_test

import (
	service_downloadfile_test "gin-play/services/downloadfile_test"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DownloadFileTestHandlerImplementation struct {
	mainService service_downloadfile_test.DownloadFileTestService
}

func (h *DownloadFileTestHandlerImplementation) Get(c *gin.Context) {
	file, err := h.mainService.GetFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Disposition", "attachment; filename=\"file.txt\"")
	c.Data(http.StatusOK, "application/octet-stream", file.FileByte)
}

func NewDownloadFileTestHandlerImplementation(mainService service_downloadfile_test.DownloadFileTestService) DownloadFileTestHandler {
	return &DownloadFileTestHandlerImplementation{mainService: mainService}
}
