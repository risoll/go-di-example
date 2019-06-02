package producthandler

import (
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/julienschmidt/httprouter"

	"github.com/risoll/tokosidia/core/product/productctrl"
	"github.com/risoll/tokosidia/utils/httputil"
)

type (
	ProductHandler interface {
		Get(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	}

	productHandlerImpl struct {
		ctrl productctrl.ProductCtrl
	}
)

func New(ctrl productctrl.ProductCtrl) ProductHandler {
	return &productHandlerImpl{
		ctrl: ctrl,
	}
}

// Get is ...
func (h *productHandlerImpl) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res := httputil.NewResponse()
	productID, _ := strconv.ParseInt(p.ByName("id"), 10, 64)
	data, err := h.ctrl.GetByID(productID)
	if err != nil {
		logrus.Errorf("[ProductHandler][Get]%s", err.Error())
		res.WriteError(w, http.StatusBadRequest, []string{"invalid product id"})
		return
	}

	res.WriteOK(w, data)
	return
}
