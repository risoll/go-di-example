package producthandler

import (
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Get is ...
func Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productID := p.ByName("id")
	logrus.Info(productID)
	return
}