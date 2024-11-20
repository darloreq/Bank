package main

import (
	"coolBank/internal/handlers"
	"coolBank/internal/infrastructure/cache"
	"coolBank/internal/infrastructure/postgres"
	bank "coolBank/internal/services/bank"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	dataBase, err := postgres.New()
	if err != nil {
		log.Fatal(err)
	}

	bankCache := cache.New()
	bankService := bank.New(dataBase)
	handler := handlers.New(bankService)
	r := chi.NewRouter()
	r.Get("/{UserID}", handler.ShowBalance)
	r.Put("/{UserID}", handler.PutMoneyIn)
	r.Post("/signIn", handler.NewUser) //СОЗДАНИЕ НОВОГО ЮЗЕРА
	http.ListenAndServe(":8080", r)
}
