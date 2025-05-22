package health_check

import (
	"github.com/arifbugaresa/go-starter/utils/response"
	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	router.GET("", func(ctx *gin.Context) {
		HealthCheckEndpoint(ctx)
	})
}

func HealthCheckEndpoint(ctx *gin.Context) {
	response.GenerateSuccessResponse(ctx, "success health check")
}
