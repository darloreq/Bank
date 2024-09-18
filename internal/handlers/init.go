package handlers

import "coolBank/internal/services/bank"

type handler struct {
	srv *bank.Bank
}

func New() *handler {
	return &handler{}
}

//прописать реализацию, все 3 метода БЛ

type HeadHandler interface {
	ShowBalance()
	PutMoneyIn()
	TakeMoneyFrom()
}
