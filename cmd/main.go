package main

import (
	"coolBank/internal/entity"
	"coolBank/internal/infrastructure/cache"
	bank "coolBank/internal/services/bank"
)

func main() {
	bankCache := cache.New()
	bankService := bank.New(bankCache)
	bankService.ShowBalance(entity.User{ID: 1}) //сюда попадёт user с сервера
}
