package cache

import (
	"coolBank/internal/entity"
	"errors"
)

type cache struct {
	bankDB map[int]entity.Balance
}

func NewCache(bankDB map[int]entity.Balance) *cache {
	return &cache{bankDB: bankDB}
}

func (c *cache) PutMoneyInCache(userID int, amountPut entity.ChangeBalance) (entity.Balance, error) {
	balance, err := c.ShowBalance(userID)

	if err != nil {
		return balance, err
	}

	balance.Numbers = balance.Numbers + amountPut.Amount

	return balance, err
}

func (c *cache) TakeMoneyFromCache(userID int, amountTake entity.ChangeBalance) (entity.Balance, error) {
	balance, err := c.ShowBalance(userID)

	if err != nil {
		return balance, err
	}

	var newBalance entity.Balance
	newBalance.Numbers = balance.Numbers - amountTake.Amount

	if newBalance.Numbers < 0 {
		return (entity.Balance{}), errors.New("insufficient balance")
	}

	return newBalance, nil
}

func (c *cache) ShowBalance(userID int) (entity.Balance, error) {
	balance, ok := c.bankDB[userID]

	if !ok {
		return (entity.Balance{}), errors.New("user not found")
	}

	return balance, nil
}
