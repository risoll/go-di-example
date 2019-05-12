package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/risoll/tokosidia/producthandler"
)

func routes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/ping", ping)
	router.GET("/product/:id", producthandler.Get)
	return router
}
