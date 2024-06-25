package user

import (
	"example.com/users/internal/domain"
	"github.com/google/uuid"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func (s Service) Create(user domain.User) (id string, err error) {
	// Conectar a la base de datos SQLite.

	// Generar un UUID para el nuevo usuario.
	user.ID = uuid.New().String()

	userID, err := s.UserRepository.Insert(user)
	if err != nil {
		return "", err
	}

	return userID, nil
}
