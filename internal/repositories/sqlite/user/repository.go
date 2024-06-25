package user

import (
	"database/sql"

	"example.com/users/internal/ports"
)

type Repository struct {
	Client *sql.DB
}

func NewRepository(db *sql.DB) ports.UserRepository {
	return Repository{
		Client: db,
	}
}
