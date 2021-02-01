package v1

import "github.com/gin-gonic/gin"

func Authentication(r *gin.RouterGroup) {
	auth := r.Group("auth")
	{
		auth.GET("login", func(c *gin.Context) {
			c.String(200, "Working Jos!")
		})
	}
}
