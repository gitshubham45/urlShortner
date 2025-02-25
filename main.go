package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/urlShortner/shortner"
)

var (
	urlMap = make(map[string]string)
	mu     sync.Mutex
)

func main() {
	fmt.Println("URL shortner")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})
	})

	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	r.POST("/url", func(c *gin.Context) {
		text := c.PostForm("text") // Get form data
		if text == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Text parameter is missing"})
			return
		}

		fmt.Println("Received:", text)
		shortenedUrl := shortner.ShortenUrl(text) // Generate shortened URL

		mu.Lock()
		urlMap[shortenedUrl] = text
		mu.Unlock()

		// Render the HTML page with the shortened URL
		c.HTML(http.StatusOK, "result.html", gin.H{
			"shortened_url": "http://localhost:3000/url/" + shortenedUrl,
		})
	})

	r.GET("/url/:shortUrl", func(c *gin.Context) {
		shortUrl := c.Param("shortUrl")

		mu.Lock()
		originalUrl, exists := urlMap[shortUrl]
		mu.Unlock()

		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
			return
		}

		c.Redirect(http.StatusFound, originalUrl)
	})

	r.Run(":3000")
}
