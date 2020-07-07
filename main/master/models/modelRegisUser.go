package models

type Register struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	Phonenumber int    `json:"phonenumber"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
