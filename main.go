package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"tnc-masters/pkg/database"
	"tnc-masters/pkg/response"
	"tnc-masters/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// init database connection
	close, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	clean(close)

	// init api route
	route := routes.NewRoute()
	route.Router.Use(gin.Logger(), response.NewXResponseTimer)
	route.Test()
	route.BU()
	route.Categories()

	if err := route.Router.Run(":8888"); err != nil {
		log.Fatal(err)
	}
}

// clean is close connection when ctrl+c for develop
func clean(tearDown func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		tearDown()
		os.Exit(0)
	}()
}
