package productservice

import (
	"testing"

	"github.com/risoll/tokosidia/models"
	"github.com/risoll/tokosidia/utils/dbutil"
)

func TestGetByID(t *testing.T) {
	fakeDB, _ := dbutil.NewConnection(dbutil.DBConf{})
	tests := []struct {
		name string
		shouldErr bool
		service ProductService
		id int64
	} {
		{
			"success",
			false,
			NewFake(fakeDB, nil, models.Product{}),
			1,
		},
	}

	for _, r := range tests {
		_, err := r.service.GetByID(r.id)
		if (err != nil) != r.shouldErr {
			t.Errorf("[TestGetByID][%s] error doesn't match", r.name)
		}	
	}
}