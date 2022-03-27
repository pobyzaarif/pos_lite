package user

import (
	"github.com/pobyzaarif/pos_lite/business"
)

type (
	service struct {
		repository Repository
	}

	Service interface {
		FindByIDandVersion(ic business.InternalContext, id int, version int) (User, error)

		FindByEmail(ic business.InternalContext, email string) (User, error)
	}
)

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) FindByIDandVersion(ic business.InternalContext, id int, version int) (User, error) {
	return s.repository.FindByIDandVersion(ic, id, version)
}

func (s *service) FindByEmail(ic business.InternalContext, email string) (User, error) {
	return s.repository.FindByEmail(ic, email)
}
