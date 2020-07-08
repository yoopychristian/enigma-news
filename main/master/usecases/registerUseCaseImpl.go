package usecases

import (
	"enigma-news/main/master/models"
	"enigma-news/main/master/repositories"
	"enigma-news/main/master/utils"
	"log"
)

type RegisterUsecaseImpl struct {
	registerRepo repositories.RegisterRepository
}

func (s RegisterUsecaseImpl) GetAllRegisters() ([]*models.Register, error) {
	register, err := s.registerRepo.GetAllRegisters()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return register, nil
}

func (s RegisterUsecaseImpl) GetRegistersID(Username string) (*models.Register, error) {
	register, err := s.registerRepo.GetRegistersID(Username)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return register, nil
}

func (s RegisterUsecaseImpl) CreateRegisters(register models.Register) error {
	err := utils.ValidateInputNotNil(register.Id, register.Email, register.Fullname, register.Phonenumber, register.Username, register.Password)
	if err != nil {
		log.Println(err)
		log.Println(err)
		return err
	}

	err = s.registerRepo.CreateRegisters(register)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s RegisterUsecaseImpl) UpdateRegisters(register models.Register) error {
	err := utils.ValidateInputNotNil(register.Id, register.Email, register.Fullname, register.Phonenumber, register.Username, register.Password)
	if err != nil {
		log.Println(err)
		return err
	}

	err = s.registerRepo.UpdateRegisters(register)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (s RegisterUsecaseImpl) DeleteRegisters(Username string) error {
	err := s.registerRepo.DeleteRegisters(Username)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func InitRegisterUseCase(registerRepo repositories.RegisterRepository) RegisterUseCase {
	return &RegisterUsecaseImpl{registerRepo}
}
