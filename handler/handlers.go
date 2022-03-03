package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resyon/url-shortener/shortener"
	"github.com/resyon/url-shortener/store"
)

// TODO: move to configuration file
const host = "http://localhost:9808/"

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationReq UrlCreationRequest	
	if err := c.ShouldBindJSON(&creationReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	shortUrl := shortener.GenerateShortLink(creationReq.LongUrl, creationReq.UserId)
	store.SaveUrlMapping(shortUrl, creationReq.LongUrl, creationReq.UserId)

	c.JSON(http.StatusOK, gin.H{
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	url := c.Param("short-url")	
	c.Redirect(http.StatusFound, store.RetrieveInitialUrl(url))
}