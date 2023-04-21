package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dylantic/SSLurper/models"
	"github.com/gin-gonic/gin"
)

func Query(ctx *gin.Context) {
	//input, error := io.ReadAll(ctx.Request.Body)
	input := ctx.Request.Body

	var jsondata models.QueryRequest
	requestDecoder := json.NewDecoder(input)
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

	//var jsondata map[string]interface{}

	ctx.JSON(http.StatusOK, jsondata)
}
