package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/risoll/tokosidia/core/product/productctrl"
	"github.com/risoll/tokosidia/core/product/producthandler"
	"github.com/risoll/tokosidia/core/product/productservice"
	"github.com/risoll/tokosidia/utils/dbutil"
	"github.com/risoll/tokosidia/utils/httputil"
	"github.com/risoll/tokosidia/utils/configutil"
)

func main() {
	
	// initiate config
	conf := configutil.Init()
	
	// initiate db
	db, err := dbutil.NewConnection(conf.DB)
	if err != nil {
		logrus.Panic("failed to connect to db", err)
		return
	}

	// initiate core module - product
	productService := productservice.New(db)
	productCtrl := productctrl.New(productService)
	productHandler := producthandler.New(productCtrl)

	// initiate instance
	port := fmt.Sprintf(":%d", conf.Instance.Port)
	router := routes(productHandler)
	server := http.Server{Handler: router, Addr: port}
	instance := httputil.NewInstance(&server)

	// start instance
	err = instance.Start()
	if err != nil {
		logrus.Panic("failed to start the instance")
		return
	}
}
