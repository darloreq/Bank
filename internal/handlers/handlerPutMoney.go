package handlers

import (
	"coolBank/internal/entity"
	"encoding/json"
	"github.com/go-chi/chi"
	"io"
	"net/http"
	"strconv"
)

type amount struct {
	TotalChange   float64 `json:"total_change"`
	OperationType string  `json:"operation_type"`
}

func (h *handler) PutMoneyIn(w http.ResponseWriter, r *http.Request) {
	UserID := chi.URLParam(r, "UserID")

	trueUserID, err := strconv.Atoi(UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var amount amount

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(body, &amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	newBalance, err := h.bankService.ChangeBalance(trueUserID, entity.ChangeBalance{Amount: amount.TotalChange}, amount.OperationType)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	jBalance, err := json.Marshal(newBalance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jBalance)
}
