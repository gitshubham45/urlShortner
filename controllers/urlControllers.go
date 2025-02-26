package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/urlShortner/db"
	"github.com/gitshubham45/urlShortner/hash"
	"github.com/gitshubham45/urlShortner/models"
	"go.mongodb.org/mongo-driver/bson"
)

func URLShortner() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.PostForm("text")
		if url == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Text parameter is missing"})
			return
		}

		fmt.Println("Received:", url)

		collection := db.OpenCollection(db.Client, "urls")

		var existingUrl models.Url
		err := collection.FindOne(c, bson.M{"longUrl": url}).Decode(&existingUrl)
		if err == nil { 
			c.HTML(http.StatusOK, "result.html", gin.H{
				"shortened_url": "http://localhost:3000/" + *existingUrl.ShortUrl,
			})
			return
		}

		shortenedUrl := hash.HashString(url)
		newUrl := models.NewUrl(&url, &shortenedUrl)

		_, err = collection.InsertOne(c, newUrl)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			return
		}

		c.HTML(http.StatusOK, "result.html", gin.H{
			"shortened_url": "http://localhost:3000/" + shortenedUrl,
		})
	}
}

func RedirectToOriginalURL() gin.HandlerFunc {
	return func(c *gin.Context) {
		shortUrl := c.Param("shortUrl")

		var url models.Url
		collection := db.OpenCollection(db.Client, "urls")
		err := collection.FindOne(c, bson.M{"shortUrl": shortUrl}).Decode(&url)
		if err != nil {
			log.Println("Shortened URL not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "Shortened URL not found"})
			return
		}

		_, err = collection.UpdateOne(
			c,
			bson.M{"shortUrl": shortUrl},
			bson.M{"$inc": bson.M{"visitCount": 1}},
		)

		if err != nil {
			log.Println("Failed to update visit count")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		urlString := url.LongUrl
		log.Println(urlString)

		c.Redirect(http.StatusFound, *urlString)
	}
}
