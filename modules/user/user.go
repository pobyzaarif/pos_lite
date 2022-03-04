package user

import (
	"github.com/pobyzaarif/pos_lite/business/user"

	"gorm.io/gorm"
)

type (
	GormRepository struct {
		db *gorm.DB
	}
)

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db.Table("users"),
	}
}

func (repo *GormRepository) FindByIDandVersion(id, version int) (selectedUser user.User, err error) {
	if err := repo.db.Select("*").Where("id = ? AND version = ?", id, version).Find((&selectedUser)).Error; err != nil {
		return selectedUser, err
	}

	return
}

func (repo *GormRepository) FindByEmail(email string) (selectedUser user.User, err error) {
	if err := repo.db.Debug().Where("email = ?", email).Find(&selectedUser).Error; err != nil {
		return selectedUser, err
	}

	return
}
