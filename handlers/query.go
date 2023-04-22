package handlers

import (
	"net/http"

	"fmt"

	"github.com/bytedance/sonic"
	"github.com/dylantic/SSLurper/models/errmsg"
	models "github.com/dylantic/SSLurper/models/queryreq"
	"github.com/gin-gonic/gin"
)

func Query(ctx *gin.Context) {
	input := ctx.Request.Body

	gin.EnableJsonDecoderDisallowUnknownFields()

	var jsondata models.QueryReq
	requestDecoder := sonic.ConfigFastest.NewDecoder(input)
	requestDecoder.DisallowUnknownFields()

	if err := requestDecoder.Decode(&jsondata); err != nil {
		resErr := errmsg.Create(1, err.Error())
		ctx.Error(err)
		fmt.Println(resErr)
		ctx.JSON(http.StatusBadRequest, resErr)
		return
	}

	ctx.JSON(http.StatusOK, jsondata)
}
