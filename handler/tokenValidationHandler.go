package handler

import (
	"enigma-news/tools"
	"fmt"
	"net/http"
)

type tokenValidationHandler struct {
}

func NewTokenValidationHandler() IHandler {
	return &tokenValidationHandler{}
}

func (h *tokenValidationHandler) Handler(w http.ResponseWriter, r *http.Request) {
	tokenValue := r.URL.Query()["token"]
	if len(tokenValue) == 0 {
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
	}

	token := tokenValue[0]

	_, err := tools.JwtDecoder(token)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
	}

	w.Write([]byte("ok"))
}
