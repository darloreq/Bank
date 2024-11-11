package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *handler) NewUser(w http.ResponseWriter, r *http.Request) {
	u := h.bankService.CreateUser()

	jUser, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jUser)
}
