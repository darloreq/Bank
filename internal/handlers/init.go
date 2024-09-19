package handlers

import (
	"coolBank/internal/entity"
)

type handler struct {
	bankService HeadHandler
}

func New(bankService HeadHandler) *handler {
	return &handler{bankService: bankService}
}

type HeadHandler interface { //методы БЛ
	ShowBalance(user entity.User) (entity.Balance, error)
	PutMoneyIn(user entity.User, amountPut entity.ChangeBalance) (entity.Balance, error)
	TakeMoneyFrom(user entity.User, amountTake entity.ChangeBalance) (entity.Balance, error)
}
