package postgres

import (
	"coolBank/internal/entity"
)

type db struct {
}

func (d *db) ShowBalance(userID int) (entity.Balance, error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) PutMoneyInCache(userID int, amountPut entity.ChangeBalance) (entity.Balance, error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) TakeMoneyFromCache(userID int, amountTake entity.ChangeBalance) (entity.Balance, error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) MakeUser() entity.User {
	//TODO implement me
	panic("implement me")
}

func New() *db {
	return &db{}
}
