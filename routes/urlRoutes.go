package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/urlShortner/controllers"
)

func UrlRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/url",controllers.URLShortner())
	incomingRoutes.GET("/:shortUrl", controllers.RedirectToOriginalURL())
}
