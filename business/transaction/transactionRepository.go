package transaction

type Repository interface {
	GetSummaryTransaction(date string) (summaryTransaction TransactionSummaryByDate, err error)
}
