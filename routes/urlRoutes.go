package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/urlShortner/controllers"
	"github.com/gitshubham45/urlShortner/cache"
	"github.com/go-redis/redis/v8"

)

var redisClient *redis.Client = cache.RedisInit()

func UrlRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/url",controllers.URLShortner(redisClient))
	incomingRoutes.GET("/:shortUrl", controllers.RedirectToOriginalURL(redisClient))
}
