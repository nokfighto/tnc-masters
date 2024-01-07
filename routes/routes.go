package routes

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Router *gin.Engine
}

func NewRoute() *Route {
	router := gin.New()
	router.HandleMethodNotAllowed = false
	//set middleware

	return &Route{
		Router: router,
	}
}

func (r *Route) Test() {
	r.Router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})
}
