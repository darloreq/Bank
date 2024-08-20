package cache

import (
	"coolBank/internal/entity"
	"errors"
	"sync"
)

type cache struct {
	bankDB map[int]entity.Balance
	mu     *sync.Mutex
}

func New() *cache {
	m := make(map[int]entity.Balance)
	mu := &sync.Mutex{}
	return &cache{bankDB: m, mu: mu}
}

func (c *cache) PutMoneyInCache(userID int, amountPut entity.ChangeBalance) (entity.Balance, error) { //TODO написать получение баланса напрямую :(

	balance, err := c.ShowBalance(userID)
	if err != nil {
		return balance, err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	balance.Numbers = balance.Numbers + amountPut.Amount

	c.bankDB[userID] = balance

	return balance, err
}

func (c *cache) TakeMoneyFromCache(userID int, amountTake entity.ChangeBalance) (entity.Balance, error) { //TODO написать получение баланса напрямую :)
	balance, err := c.ShowBalance(userID)
	if err != nil {
		return balance, err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	var newBalance entity.Balance
	newBalance.Numbers = balance.Numbers - amountTake.Amount
	if newBalance.Numbers < 0 {
		return (entity.Balance{}), errors.New("insufficient balance")
	}

	c.bankDB[userID] = balance

	return newBalance, nil
}

func (c *cache) ShowBalance(userID int) (entity.Balance, error) {
	c.mu.Lock()

	defer c.mu.Unlock()

	balance, ok := c.bankDB[userID]
	if !ok {
		return (entity.Balance{}), errors.New("user not found")
	}

	return balance, nil
}
