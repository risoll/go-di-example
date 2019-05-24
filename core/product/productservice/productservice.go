package productservice

import (
	"github.com/risoll/tokosidia/models"
	"github.com/risoll/tokosidia/utils/dbutil"
)

type (
	ProductService interface {
		GetByID(id int64) (models.Product, error)
	}

	productServiceImpl struct {
		db dbutil.DB
	}
)

func New(db dbutil.DB) ProductService {
	return &productServiceImpl{db: db}
}

func (s *productServiceImpl) GetByID(id int64) (models.Product, error){
	query := `SELECT * from product where id = $1`
	var result models.Product
	err := s.db.Select(&result, query, id)
	return result, err
}
