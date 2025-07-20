package bank

import (
	"coolBank/internal/entity"
)

func (b *Bank) CreateUser(user entity.CreateUser) (entity.User, error) {

	newUser, err := b.repos.MakeUser(user)
	if err != nil {
		return entity.User{}, err
	}
	return newUser, err
}
