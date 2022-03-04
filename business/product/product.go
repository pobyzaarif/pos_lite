package product

import "github.com/pobyzaarif/pos_lite/business"

type (
	Product struct {
		ID    int
		Name  string
		Price int

		business.ObjectMetadata
	}
)
