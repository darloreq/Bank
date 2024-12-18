package bank

import "coolBank/internal/entity"

func (b *Bank) PutMoneyIn(user entity.User, amountPut entity.ChangeBalance) (entity.Balance, error) {
	balance, err := b.ShowBalance(user)
	if err != nil {
		return balance, err
	}

	var modifiedBalance entity.ChangeBalance
	modifiedBalance.Amount = balance.Numbers + amountPut.Amount

	balanceFromDB, err := b.repos.PutMoneyInDB(user.ID, modifiedBalance) //запись в КЕШ
	if err != nil {
		return balance, err
	}

	return balanceFromDB, nil
}
