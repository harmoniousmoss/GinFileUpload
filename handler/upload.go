package handler

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	fmt.Println("GOOGLE_APPLICATION_CREDENTIALS:", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Error receiving file: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	// Open the file
	f, err := file.Open()
	if err != nil {
		log.Printf("Error opening file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the file"})
		return
	}
	defer f.Close()

	// Initialize Google Cloud Storage client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("Error creating storage client: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create storage client"})
		return
	}

	bucketName := os.Getenv("GCS_BUCKET")
	if bucketName == "" {
		log.Println("GCS_BUCKET environment variable is not set")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "GCS bucket name is not set"})
		return
	}

	bucket := client.Bucket(bucketName)

	// Upload the file
	obj := bucket.Object(file.Filename)
	w := obj.NewWriter(ctx)
	defer w.Close()

	if _, err := io.Copy(w, f); err != nil {
		log.Printf("Error writing file to storage: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to write the file to storage"})
		return
	}

	if err := w.Close(); err != nil {
		log.Printf("Error closing storage writer: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to close the storage writer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "filename": file.Filename})
}
