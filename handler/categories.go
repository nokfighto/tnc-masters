package handler

import (
	"net/http"
	"tnc-masters/pkg/response"
	"tnc-masters/services/categories"

	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	data, err := categories.List("", "Y")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.ResponseSuccess(data))
}
