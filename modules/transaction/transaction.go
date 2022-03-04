package transaction

import (
	"github.com/pobyzaarif/pos_lite/business/transaction"

	"gorm.io/gorm"
)

type (
	GormRepository struct {
		db *gorm.DB
	}
)

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db.Table("transactions"),
	}
}

func (repo *GormRepository) GetSummaryTransaction(date string) (summaryTransaction transaction.TransactionSummaryByDate, err error) {
	dateStart := date + " 00:00:01"
	dateEnd := date + " 23:59:59"
	if err := repo.db.Select("COUNT(*) total_transaction, SUM(qty) total_products_sold, SUM(total_price) total_gross_earnings").Where("created_at BETWEEN ? AND ?", dateStart, dateEnd).Find((&summaryTransaction)).Error; err != nil {
		return summaryTransaction, err
	}

	return
}
