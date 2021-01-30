package v1

import (
	"github.com/basic/account-service/controllers"
	"github.com/gin-gonic/gin"
)

func RouterAuthentication(routingGroup *gin.RouterGroup) {
	routingGroup.GET("/hello")
}