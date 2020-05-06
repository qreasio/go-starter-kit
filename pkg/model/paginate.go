package model

// Pagination to represents common parameters for endpoint that implement pagination
type Pagination struct {
	Page  int    `validate:"gt=0,lte=10000"`
	Limit int    `validate:"gt=0,lte=10000"`
	Sort  string `validate:"oneof=asc desc"`
}

// NewPagination returns new Pagination struct with default values
func NewPagination() *Pagination {
	return &Pagination{
		Page:  1,
		Limit: 2,
		Sort:  "asc",
	}
}
