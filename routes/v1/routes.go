package v1

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	v1 := router.Group("v1")
	{

		Authentication(v1)
	}
}
