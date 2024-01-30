package routes

import "tnc-masters/handler/role"

func (r *Route) Role() {
	route := r.Router.Group("/role")
	route.GET("/list", role.List)
}
