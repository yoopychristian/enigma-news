package handler

import (
	"encoding/json"
	"enigma-news/main/master/models"
	"log"
	"net/http"

	"github.com/edwardsuwirya/mdw/utils"
)

type authenticationHandler struct {
}

func NewAuthenticationHandler() IHandler {
	return &authenticationHandler{}
}

func (h *authenticationHandler) Handler(w http.ResponseWriter, r *http.Request) {
	var sysUser models.Register
	json.NewDecoder(r.Body).Decode(&sysUser)

	userName := sysUser.Username
	userPassword := sysUser.Password

	if userName == "alexcuy" && userPassword == "pwpwpw" {
		w.WriteHeader(http.StatusOK)
		token, err := utils.JwtEncoder(userName, "Rahasia dong")
		if err != nil {
			log.Print(err)
			http.Error(w, "Failed token generation", http.StatusInternalServerError)
		}
		w.Write([]byte(token))
	} else {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
	}

}
