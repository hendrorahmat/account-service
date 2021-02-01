package v1

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	v1 := router.Group("v1")
	{

		v1.GET("/oauth/login", func(c *gin.Context) {
			c.String(200, "Working!")
		})
	}
}
