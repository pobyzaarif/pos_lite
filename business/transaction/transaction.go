package transaction

import "github.com/pobyzaarif/pos_lite/business"

type (
	Transaction struct {
		ID        int
		ProductID int
		Qty       int
		Price     int
		Total     int

		business.ObjectMetadata
	}

	TransactionSummaryByDate struct {
		TotalTransaction   int `json:"total_transaction"`
		TotalProductsSold  int `json:"total_products_sold"`
		TotalGrossEarnings int `json:"total_gross_earnings"`
	}
)
