package routes

import (
	"tnc-masters/handler"
)

func (r *Route) Categories() {
	x := r.Router.Group("cate")
	x.GET("/list", handler.List)
}
