package bank

import "coolBank/internal/entity"

func (b *Bank) CreateUser(user entity.CreateUser) entity.User {

	newUser, err := b.repos.MakeUser(user)
	if err != nil {
		return entity.User{}
	}
	return newUser
}
