package main

import (
	"gingcs/route" // Replace with your project's module name

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()

	// Setup the router
	r := route.SetupRouter()

	// Start the server
	r.Run() // listen and serve on 0.0.0.0:8080
}
