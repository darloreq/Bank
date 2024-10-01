package handlers

import (
	"coolBank/internal/entity"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

// что значит запись (h *handler)
func (h *handler) PutMoneyIn(w http.ResponseWriter, r *http.Request) {
	UserID := chi.URLParam(r, "UserID")

	trueUserID, err := strconv.Atoi(UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//тут мы должны получить сумму, на которую пополняют баланс. Как это будет выглядеть в интерфейсе? Как это получить

	balance, err := h.bankService.ShowBalance(entity.User{ID: trueUserID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//тут изменение бала

}
