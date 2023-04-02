package repository

type Repository struct {
	User IUserRepository
}

func New() *Repository {
	user := NewUser()
	return &Repository{
		User: user,
	}
}
