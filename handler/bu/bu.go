package bu

import (
	"log"
	"net/http"
	"tnc-masters/pkg/response"
	"tnc-masters/services/bu"

	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	data, err := bu.List()
	if err != nil {
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, response.ResponseSuccess(data))
}
