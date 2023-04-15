package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()
	// A GET TEST ROUTE
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{"hello": "got the request"})
	})
	// Define a route to handle file uploads
	r.POST("/upload", func(c *gin.Context) {
		// Get the file from the request
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Open the file
		f, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer f.Close()

		// Generate the MD5 checksum from the file contents
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Convert the checksum to a string and return it
		sum := fmt.Sprintf("%x", h.Sum(nil))
		c.JSON(http.StatusOK, gin.H{"checksum": sum})
	})

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
