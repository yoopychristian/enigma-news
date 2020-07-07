package controllers

import (
	"encoding/json"
	"enigma-news/main/master/models"
	"enigma-news/main/master/response"
	"enigma-news/main/master/usecases"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RegisterHandler struct {
	registerUseCase usecases.RegisterUseCase
}

func RegisterController(r *mux.Router, service usecases.RegisterUseCase) {
	RegisterHandler := RegisterHandler{service}
	r.HandleFunc("/registers", RegisterHandler.ListRegisters).Methods(http.MethodGet)
	r.HandleFunc("/register/{Username}", RegisterHandler.GetRegistersID).Methods(http.MethodGet)
	r.HandleFunc("/register", RegisterHandler.CreateDataRegisters).Methods(http.MethodPost)
	r.HandleFunc("/register", RegisterHandler.UpdateDataRegisters).Methods(http.MethodPut)
	r.HandleFunc("/register/{Username}", RegisterHandler.DeleteDataRegisters).Methods(http.MethodDelete)
}

func (s RegisterHandler) ListRegisters(w http.ResponseWriter, r *http.Request) {
	register, err := s.registerUseCase.GetAllRegisters()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Get Database Registers"
	pesan.Data = register

	byteOfRegisters, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfRegisters)
	fmt.Println("EndPoint Hit : GetRegisterPage")
}

func (s RegisterHandler) GetRegistersID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	register, err := s.registerUseCase.GetRegistersID(vars["Username"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Get Database Register by Username"
	pesan.Data = register

	byteOfRegisters, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfRegisters)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: GetRegisterByIDPage")
}

func (s RegisterHandler) CreateDataRegisters(w http.ResponseWriter, r *http.Request) {
	var register models.Register
	_ = json.NewDecoder(r.Body).Decode(&register) // json ke struct
	s.registerUseCase.CreateRegisters(register)
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Create Data for Register Database"
	pesan.Data = "Success"

	byteOfRegisters, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfRegisters)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: CreateRegisterPage")
}

func (s RegisterHandler) UpdateDataRegisters(w http.ResponseWriter, r *http.Request) {
	var register models.Register
	_ = json.NewDecoder(r.Body).Decode(&register) // json ke struct
	s.registerUseCase.UpdateRegisters(register)

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Update Data for Register Database"
	pesan.Data = "Success"

	byteOfRegisters, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfRegisters)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: UpdateRegisterPage")
}

func (s RegisterHandler) DeleteDataRegisters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s.registerUseCase.DeleteRegisters(vars["Username"])
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Delete Database Register"
	pesan.Data = "Success"

	byteOfRegisters, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfRegisters)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: DeleteRegisterPage")
}
