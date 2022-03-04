package user

type Repository interface {
	FindByIDandVersion(id, version int) (selectedUser User, err error)

	FindByEmail(email string) (selectedUser User, err error)
}
