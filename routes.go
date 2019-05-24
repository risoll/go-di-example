package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/risoll/tokosidia/core/product/producthandler"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "pong")
}

func routes(productHandler producthandler.ProductHandler) *httprouter.Router {
	router := httprouter.New()
	router.GET("/ping", ping)
	router.GET("/product/:id", productHandler.Get)
	return router
}
