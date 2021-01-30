package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hendrorahmat/account-service/controllers"
)

func RouterAuthentication(routingGroup *gin.RouterGroup) {
	routingGroup.GET("/hello")
}
