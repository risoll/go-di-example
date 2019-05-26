package dbutil

type (
	resultFake struct {

	}
)

func (r resultFake) LastInsertId() (int64, error) {
	return r.RowsAffected()
}

func (r resultFake) RowsAffected() (int64, error) {
	return 0, nil
}