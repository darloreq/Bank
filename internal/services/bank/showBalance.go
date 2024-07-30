package bank

import "coolBank/internal/entity"

func (b *Bank) ShowBalance(user entity.User) (entity.Balance, error) {
	balance, err := b.Repos.ShowBalance(user.ID)

	if err != nil {
		return balance, err
	}

	return balance, nil
}