package routes

import (
	"shorturlsrv/app/controllers"
	"shorturlsrv/app/middleware"

	"github.com/gin-gonic/gin"
	// https://91porn.com/view_video.php?viewkey=e96a0531fa736f0c3e3d&page=&viewtype=&category=
)

func Api(router *gin.Engine) {
	apiGroup := router.Group("api")

	shortUrlGroup := apiGroup.Group("/short_urls")
	shortUrlGroup.Use(middleware.CheckSecret())
	shortUrlGroup.GET("", controllers.ShortUrlController.List)
	shortUrlGroup.GET("/:id", controllers.ShortUrlController.Show)
	shortUrlGroup.POST("/", controllers.ShortUrlController.Store)
	shortUrlGroup.PUT("/:id", controllers.ShortUrlController.Update)
	shortUrlGroup.DELETE("/:id", controllers.ShortUrlController.Destroy)

	authGroup := apiGroup.Group("/auth")
	authGroup.POST("/register", controllers.AuthController.Register)
	authGroup.POST("/login", controllers.AuthController.Login)
}
