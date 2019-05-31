package productctrl

import (
	"github.com/risoll/tokosidia/core/product/productservice"
	"github.com/risoll/tokosidia/models"
)

type (
	ProductCtrl interface {
		GetByID(id int64) (models.Product, error)
	}

	// TODO add more dependencies
	productCtrlImpl struct {
		service productservice.ProductService
	}
)

func New(service productservice.ProductService) ProductCtrl {
	return &productCtrlImpl{service: service}
}

func (c *productCtrlImpl) GetByID(id int64) (models.Product, error) {
	return c.service.GetByID(id)
}