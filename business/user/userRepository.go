package user

import (
	"github.com/pobyzaarif/pos_lite/business"
)

type Repository interface {
	FindByIDandVersion(ic business.InternalContext, id, version int) (selectedUser User, err error)

	FindByEmail(ic business.InternalContext, email string) (selectedUser User, err error)
}
