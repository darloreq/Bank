package bank

import "coolBank/internal/entity"

func (b *Bank) PutMoneyIn(user entity.User, amount entity.ChangeBalance) (entity.Balance, error) {
	balance, err := b.ShowBalance(user)
	if err != nil {
		return balance, err
	}

	var modifiedBalance entity.ChangeBalance
	modifiedBalance.Amount = balance.Numbers + amount.Amount

	balanceFromCache, err := b.Repos.PutMoneyInCache(user.ID, modifiedBalance) //запись в КЕШ
	if err != nil {
		return balance, err
	}
	//z goida
	return balanceFromCache, nil
}
