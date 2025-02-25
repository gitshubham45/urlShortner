package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/urlShortner/routes"
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

	routes.UrlRoutes(r)

	r.Run(":3000")
}
