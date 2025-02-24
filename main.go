package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/urlShortner/shortner"
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
		text := c.PostForm("text") // Correct way to get form data
		if text == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Text parameter is missing"})
			return
		}

		fmt.Println("Received:", text)
		shortenedUrl := shortner.ShortenUrl(text) // Assuming this function exists

		c.JSON(http.StatusOK, gin.H{"message": "Success", "shortened_url": shortenedUrl})
	})

	r.Run(":3000")
}
