package handlers

import (
	"coolBank/internal/entity"
)

//go:generate mockgen -source=init.go -destination=mocks/mock.go

type handler struct {
	bankService HeadHandler
}

func New(bankService HeadHandler) *handler {
	return &handler{bankService: bankService}
}

type HeadHandler interface { //методы БЛ
	ShowBalance(user entity.User) (entity.Balance, error)
	ChangeBalance(userID int, amount entity.ChangeBalance, operationType string) (entity.Balance, error)
	CreateUser(user entity.CreateUser) entity.User
}
