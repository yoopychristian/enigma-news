package repositories

import "enigma-news/main/master/models"

type UserRepository interface {
	Get(*models.User) (bool, error)
}
