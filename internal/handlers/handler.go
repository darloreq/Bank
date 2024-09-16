package handlers

import "net/http"

type handler struct {
}

func New() *handler {
	return &handler{}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
