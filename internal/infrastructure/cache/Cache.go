package cache

import (
	"coolBank/internal/entity"
)

type cache struct {
	bankDB map[int]entity.Balance
}

func NewCache(bankDB map[int]entity.Balance) *cache {
	return &cache{bankDB: bankDB}
}

func (c *cache) PutMoneyInCache(userID int, amountPut entity.ChangeBalance) (entity.Balance, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cache) TakeMoneyFromCache(userID int, amountTake entity.ChangeBalance) (entity.Balance, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cache) ShowBalance(userID int) (entity.Balance, error) {
	panic("implement me")
}
