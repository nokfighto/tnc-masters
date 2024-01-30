package role

import (
	"log"
	"net/http"
	"tnc-masters/pkg/response"
	"tnc-masters/services/role"

	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	data, err := role.List()
	if err != nil {
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, response.ResponseSuccess(data))
}
