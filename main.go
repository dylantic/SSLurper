package main

import (
	"fmt"
	"time"

	"github.com/dylantic/SSLurper/handlers"
	"github.com/dylantic/SSLurper/helpers"
	"github.com/dylantic/SSLurper/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	Config := &helpers.Config

	if !Config.Server.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Use(middleware.AddReqID())

	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - %s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.Keys["ReqID"],
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}))

	r.GET("/test", handlers.Test)

	r.POST("/query", handlers.Query)

	r.Run(":8080")

}
