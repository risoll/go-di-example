package main

import (
	"fmt"
	"github.com/risoll/tokosidia/constants"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "pong")
}

func main() {
	router := routes()
	server := http.Server{Handler:router, Addr: constants.Port}
	instance := newInstance(&server)
	err := instance.Start()
	if err != nil {
		logrus.Panic("failed to start the instance")
	}
}
