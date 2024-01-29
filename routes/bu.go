package routes

import (
	"tnc-masters/handler/bu"
)

func (r *Route) BU() {
	b := r.Router.Group("/bu")
	b.GET("/list", bu.List)
}
