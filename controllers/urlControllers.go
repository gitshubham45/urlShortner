package controllers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/urlShortner/hash"
)

var (
	urlMap = make(map[string]string)
	mu     sync.Mutex
)

func URLShortner() gin.HandlerFunc {
	return func(c *gin.Context) {
		text := c.PostForm("text") // Get form data
		if text == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Text parameter is missing"})
			return
		}

		fmt.Println("Received:", text)
		shortenedUrl := hash.HashString(text) // Generate shortened URL

		mu.Lock()
		urlMap[shortenedUrl] = text
		mu.Unlock()

		// Render the HTML page with the shortened URL
		c.HTML(http.StatusOK, "result.html", gin.H{
			"shortened_url": "http://localhost:3000/url/" + shortenedUrl,
		})
	}
}

func RedirectToOriginalURL() gin.HandlerFunc {
	return func(c *gin.Context) {
		shortUrl := c.Param("shortUrl")

		mu.Lock()
		originalUrl, exists := urlMap[shortUrl]
		mu.Unlock()

		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
			return
		}

		c.Redirect(http.StatusFound, originalUrl)
	}
}
