package ports

import "example.com/users/internal/domain"

type UserService interface {
	Create(user domain.User) (id string, err error)
}

type UserRepository interface {
	Insert(user domain.User) (id string, err error)
}
