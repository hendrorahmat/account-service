package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vsouza/go-gin-boilerplate/controllers"
	"github.com/vsouza/go-gin-boilerplate/middlewares"
	RouteV1 "github.com/vsouza/go-gin-boilerplate/routes/v1"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	router.Use(middlewares.AuthMiddleware())
	RouteV1.Routes(router)
	// router.Use(middlewares.AuthMiddleware())
	// RouterV1(router)
	// v1 := router.Group("v1")
	// {
	// 	userGroup := v1.Group("user")
	// 	{
	// 		user := new(controllers.UserController)
	// 		userGroup.GET("/:id", user.Retrieve)
	// 	}
	// }
	return router

}
