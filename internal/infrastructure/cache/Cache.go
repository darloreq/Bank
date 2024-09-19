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

func New() *cache { //ДОБАВЬ СОЗДАНИЕ ЮЗЕРА и хендлер для создания юзера(изм. добавить это в методе хендлера, БЛ, кеша. Присвоение id идёт в кеше через библиотеку rand)
	m := make(map[int]entity.Balance)
	m[1] = entity.Balance{Numbers: 9011462}
	mu := &sync.Mutex{}
	return &cache{bankDB: m, mu: mu}
}

func (c *cache) PutMoneyInCache(userID int, amountPut entity.ChangeBalance) (entity.Balance, error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	balance := c.bankDB[userID]

	balance.Numbers = balance.Numbers + amountPut.Amount

	c.bankDB[userID] = balance

	return balance, nil
}

func (c *cache) TakeMoneyFromCache(userID int, amountTake entity.ChangeBalance) (entity.Balance, error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	balance := c.bankDB[userID]

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
