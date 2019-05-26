package productctrl

import (
	"testing"

	"github.com/risoll/tokosidia/core/product/productservice"
	"github.com/risoll/tokosidia/models"
	"github.com/risoll/tokosidia/utils/dbutil"
)

func TestGetByID(t *testing.T) {
	fakeDB, _ := dbutil.NewConnection(dbutil.DBConf{})
	tests := []struct {
		name      string
		shouldErr bool
		service   productservice.ProductService
		id        int64
	}{
		{
			"success",
			false,
			productservice.NewFake(fakeDB, nil, models.Product{}),
			1,
		},
	}

	for _, r := range tests {
		ctrl := New(r.service)
		_, err := ctrl.GetByID(r.id)
		if (err != nil) != r.shouldErr {
			t.Errorf("[TestGetByID][%s] error doesn't match", r.name)
		}
	}
}
