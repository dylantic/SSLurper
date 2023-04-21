package middleware

import (
	"github.com/dylantic/SSLurper/helpers"
	"github.com/gin-gonic/gin"
)

func AddReqID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := helpers.RandomBase64String(15)
		ctx.Set("ReqID", requestID)
		ctx.Header("Request-Id", requestID)
	}
}
