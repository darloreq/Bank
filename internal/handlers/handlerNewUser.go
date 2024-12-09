package handlers

import (
	"coolBank/internal/entity"
	"encoding/json"
	"io"
	"net/http"
)

func (h *handler) NewUser(w http.ResponseWriter, r *http.Request) {
	Name, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var createUser entity.CreateUser

	err = json.Unmarshal(Name, &createUser)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	u := h.bankService.CreateUser(createUser)

	jUser, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jUser)
}
