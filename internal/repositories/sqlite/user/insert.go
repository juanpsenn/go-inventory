package user

import (
	"example.com/users/internal/domain"
)

func (r Repository) Insert(user domain.User) (string, error) {
	_, err := r.Client.Exec("INSERT INTO users (id, name) VALUES (?, ?)", user.ID, user.Name)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
