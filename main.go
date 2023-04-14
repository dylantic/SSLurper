package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	r.POST("/query", func(ctx *gin.Context) {

		input, error := io.ReadAll(ctx.Request.Body)
		if error != nil {
			log.Fatalf("impossible to read all body of response: %s", error)
		}

		var jsondata map[string]interface{}

		json.Unmarshal(input, &jsondata)
		fmt.Println(jsondata)

		ctx.JSON(http.StatusOK, jsondata)
	})

	r.Run()
}
