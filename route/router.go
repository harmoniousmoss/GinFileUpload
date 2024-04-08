package route

import (
	"gingcs/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Add a GET route for the root path
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Gin-GCS Uploader")
	})

	// POST route for file upload
	r.POST("/upload", handler.UploadFile)

	return r
}
