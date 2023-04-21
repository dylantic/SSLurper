package handlers

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/dylantic/SSLurper/models"
	"github.com/gin-gonic/gin"
)

func Query(ctx *gin.Context) {
	input := ctx.Request.Body

	gin.EnableJsonDecoderDisallowUnknownFields()

	var jsondata models.QueryRequest
	requestDecoder := sonic.ConfigFastest.NewDecoder(input)
	requestDecoder.DisallowUnknownFields()

	if err := requestDecoder.Decode(&jsondata); err != nil {
		resErr := models.Error{
			Code: 1,
			Text: err.Error(),
		}
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, resErr)
		return
	}

	ctx.JSON(http.StatusOK, jsondata)
}
