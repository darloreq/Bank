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
	handler := handlers.New()
	r := chi.NewRouter()
	http.ListenAndServe(":8080", r)
}
