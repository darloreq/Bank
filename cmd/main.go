package main

import (
	"coolBank/internal/handlers"
	"coolBank/internal/infrastructure/cache"
	bank "coolBank/internal/services/bank"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	bankCache := cache.New()
	bankService := bank.New(bankCache)
	handler := handlers.New(bankService)
	r := chi.NewRouter()
	r.Get("/{UserID}", handler.ShowBalance)
	r.Put("/{UserID}", handler.PutMoneyIn)
	r.Post() //СОЗДАНИЕ НОВОГО ЮЗЕРА
	http.ListenAndServe(":8080", r)
}
