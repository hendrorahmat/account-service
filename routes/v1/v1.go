package v1

import (
	"github.com/gin-gonic/gin"
)

func RouterV1(r *gin.Engine) {
	v1 := r.Group("v1")
	{
		authentication := v1.Group("authentication")
		{
			RouterAuthentication(authentication)
		}
	}
}
