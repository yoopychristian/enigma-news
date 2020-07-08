package usecases

import "enigma-news/main/master/models"

type UserUsecase interface {
	GetUser(*models.User) (bool, error)
}
