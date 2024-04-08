package route

import (
	"gingcs/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/upload", handler.UploadFile)

	return r
}
