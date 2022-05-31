package user

import (
	"github.com/pobyzaarif/pos_lite/business"
	"github.com/pobyzaarif/pos_lite/business/user"
	"gorm.io/gorm"
)

type (
	GormRepository struct {
		*gorm.DB
	}
)

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db.Table("users"),
	}
}

func (repo *GormRepository) FindByIDandVersion(ic business.InternalContext, id, version int) (selectedUser user.User, err error) {
	query := repo.DB.WithContext(ic.ToContext())
	if err := query.Where("id = ? AND version = ?", id, version).Find((&selectedUser)).Error; err != nil {
		return selectedUser, err
	}

	return
}

func (repo *GormRepository) FindByEmail(ic business.InternalContext, email string) (selectedUser user.User, err error) {
	query := repo.DB.WithContext(ic.ToContext())
	if err := query.Where("email = ?", email).Find(&selectedUser).Error; err != nil {
		return selectedUser, err
	}

	return
}
