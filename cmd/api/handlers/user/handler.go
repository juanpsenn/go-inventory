package user

import "example.com/users/internal/ports"

type Handler struct {
	UserService ports.UserService
}