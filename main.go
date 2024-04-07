package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin engine.
	r := gin.Default()

	// Define a route and a handler function.
	r.GET("/", func(c *gin.Context) {
		// Respond with a plain text string
		c.String(200, "Hello, Gin!")
	})

	// Run the web server on port 8080.
	r.Run() // listen and serve on 0.0.0.0:8080
}
