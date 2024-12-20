package handlers

import (
	"coolBank/internal/entity"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (h *handler) ShowBalance(w http.ResponseWriter, r *http.Request) {
	UserID := chi.URLParam(r, "UserID")

	trueUserID, err := strconv.Atoi(UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	balance, err := h.bankService.ShowBalance(entity.User{ID: trueUserID})
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
}
