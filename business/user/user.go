package user

import "github.com/pobyzaarif/pos_lite/business"

type (
	User struct {
		ID       int
		Role     string
		Name     string
		Email    string
		Password string

		business.ObjectMetadata
	}
)
