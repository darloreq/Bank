package bank

import (
	"coolBank/internal/entity"
	"errors"
)

//TODO добавить разделитель операций т.к. их будет больше

type Bank struct {
	repos BankReposI
}

func New(repos BankReposI) *Bank {
	return &Bank{repos: repos}
}

type BankReposI interface {
	ShowBalance(userID int) (entity.Balance, error)
	PutMoneyInCache(userID int, amountPut entity.ChangeBalance) (entity.Balance, error)
	TakeMoneyFromCache(userID int, amountTake entity.ChangeBalance) (entity.Balance, error)
}

var (
	NoUserError      = errors.New("Такого пользователя нет в системе")
	NotEnoughBalance = errors.New("Недостаточно средств на балансе")
)
