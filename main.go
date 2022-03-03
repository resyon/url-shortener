package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/resyon/url-shortener/handler"
	"github.com/resyon/url-shortener/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})

	r.POST("/create-short-url", handler.CreateShortUrl)
	r.GET("/:short-url", handler.HandleShortUrlRedirect)

	store.InitStorageService()
	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}