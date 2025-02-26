package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/urlShortner/cache"
	"github.com/gitshubham45/urlShortner/db"
	"github.com/gitshubham45/urlShortner/hash"
	"github.com/gitshubham45/urlShortner/models"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
)

func URLShortner(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.PostForm("text")
		if url == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Text parameter is missing"})
			return
		}

		fmt.Println("Received:", url)

		// Check if URL is already cached

		cachedUrl, _ := cache.GetCachedURL(redisClient, url)
		if cachedUrl != "" {
			c.HTML(http.StatusOK, "result.html", gin.H{
				"shortened_url": "http://localhost:3000/" + cachedUrl,
			})
			return
		}

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

		cache.CacheURL(redisClient, shortenedUrl, url)

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

func RedirectToOriginalURL(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortUrl := c.Param("shortUrl")

		longUrl, err := cache.GetCachedURL(redisClient, shortUrl)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
			return
		}

		if longUrl != "" {
			log.Println("Found in cache")
		}

		if longUrl == "" {
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

			longUrl = *url.LongUrl
			cache.CacheURL(redisClient, shortUrl, longUrl)
		}

		log.Println(longUrl)

		c.Redirect(http.StatusFound, longUrl)
	}
}
