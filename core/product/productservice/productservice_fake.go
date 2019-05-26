package productservice

import (
	"github.com/risoll/tokosidia/models"
	"github.com/risoll/tokosidia/utils/dbutil"
)

type (

	// easily control error flow
	productServiceFake struct {
		db dbutil.DB
		err error
		result models.Product
	}
)

func NewFake(db dbutil.DB, err error, result models.Product) ProductService {
	return &productServiceFake{db: db, err: err, result: result}
}

func (s *productServiceFake) GetByID(id int64) (models.Product, error) {
	return s.result, s.err
}