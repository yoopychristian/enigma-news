package usecases

import "enigma-news/main/master/models"

type RegisterUseCase interface {
	GetAllRegisters() ([]*models.Register, error)
	GetRegistersID(Username string) (*models.Register, error)
	CreateRegisters(register models.Register) error
	UpdateRegisters(register models.Register) error
	DeleteRegisters(Username string) error
}
