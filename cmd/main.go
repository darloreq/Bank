package main

import (
	"coolBank/internal/entity"
	"coolBank/internal/infrastructure/cache"
	bank "coolBank/internal/services/bank"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func main() {
	bankCache := cache.New()
	bankService := bank.New(bankCache)
	r := chi.NewRouter()
	r.Get("/{UserID}", func(w http.ResponseWriter, r *http.Request) {
		UserID := chi.URLParam(r, "UserID")

		trueUserID, err := strconv.Atoi(UserID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		balance, err := bankService.ShowBalance(entity.User{ID: trueUserID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jBalance, err := json.Marshal(balance)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jBalance)
	})
	http.me
	http.ListenAndServe(":8080", r)
}
