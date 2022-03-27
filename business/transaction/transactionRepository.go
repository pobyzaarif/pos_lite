package transaction

import "github.com/pobyzaarif/pos_lite/business"

type Repository interface {
	GetSummaryTransaction(ic business.InternalContext, date string) (summaryTransaction TransactionSummaryByDate, err error)
}
