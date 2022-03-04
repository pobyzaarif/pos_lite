package user

import "github.com/pobyzaarif/pos_lite/business"

const (
	RoleSuperAdmin Role = "superadmin"
	RoleOwner      Role = "owner"
	RoleAdmin      Role = "admin"
	RoleCashier    Role = "cashier"
)

type (
	Role string

	User struct {
		ID       int
		Role     Role
		Name     string
		Email    string
		Password string

		business.ObjectMetadata
	}
)
