package route

import (
	"gingcs/handler"
	"net/http" // Added this import for http status codes

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() // Ensure this line is present to initialize the router

	// Add a GET route for the root path to serve HTML content
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome to Gin GCS Uploader</title>
</head>
<body>
    <h1>Welcome to Gin GCS Uploader</h1>
</body>
</html>`)
	})

	// POST route for file upload
	r.POST("/upload", handler.UploadFile)

	return r
}
