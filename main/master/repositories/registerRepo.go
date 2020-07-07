package repositories

import "enigma-news/main/master/models"

type RegisterRepository interface {
	GetAllRegisters() ([]*models.Register, error)
	GetRegistersID(Username string) (*models.Register, error)
	CreateRegisters(register models.Register) error
	UpdateRegisters(register models.Register) error
	DeleteRegisters(Username string) error
}
