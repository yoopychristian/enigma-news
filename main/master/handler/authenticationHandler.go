package handler

import (
	"encoding/json"
	"enigma-news/main/master/models"
	"log"
	"net/http"

	"enigma-news/main/master/tools"
)

type authenticationHandler struct {
}

func NewAuthenticationHandler() IHandler {
	return &authenticationHandler{}
}

func (h *authenticationHandler) Handler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	userName := user.Username
	userPassword := user.Password

	if userName == "alexcuy" && userPassword == "pwpwpw" {
		w.WriteHeader(http.StatusOK)
		token, err := tools.JwtEncoder(userName, "Rahasia")
		if err != nil {
			log.Print(err)
			http.Error(w, "Failed token generation", http.StatusInternalServerError)
		}
		w.Write([]byte(token))
	} else {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
	}

}
