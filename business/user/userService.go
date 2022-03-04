package user

type (
	service struct {
		repository Repository
	}

	Service interface {
		FindByIDandVersion(id int, version int) (User, error)

		FindByEmail(email string) (User, error)
	}
)

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) FindByIDandVersion(id int, version int) (User, error) {
	return s.repository.FindByIDandVersion(id, version)
}

func (s *service) FindByEmail(email string) (User, error) {
	return s.repository.FindByEmail(email)
}
