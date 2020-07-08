package usecases

import (
	"enigma-news/main/master/models"
	"enigma-news/main/master/repositories"
)

type UserUsecaseImpl struct {
	userRepo repositories.UserRepository
}

func (u *UserUsecaseImpl) GetUser(data *models.User) (bool, error) {
	isValid, err := u.userRepo.Get(data)
	if err != nil {
		return false, err
	}
	return isValid, nil
}

func InitUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}
