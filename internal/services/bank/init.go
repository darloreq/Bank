package bank

import (
	"coolBank/internal/entity"
	"errors"
)

type Bank struct {
	Repos BankReposI
}

type BankReposI interface {
	ShowBalance(userID int) (entity.Balance, error)
	PutMoneyInCache(userID int, amount entity.ChangeBalance) (entity.Balance, error)
}

var (
	NoUserError = errors.New("Такого пользователя нет в системе")
)
