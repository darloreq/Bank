package bank

import "coolBank/internal/entity"

func (b *Bank) CreateUser() entity.User {

	newUser := b.repos.MakeUser()
	return newUser

}
