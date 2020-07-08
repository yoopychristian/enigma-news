package repositories

import (
	"database/sql"
	"enigma-news/main/master/models"
)

type UserRepoImpl struct {
	db *sql.DB
}

func (u *UserRepoImpl) Get(userIn *models.User) (bool, error) {

	row := u.db.QueryRow(`select username,password from user where username=? and password =?`, userIn.Username, userIn.Password)
	var user = models.User{}
	err := row.Scan(&user.Username, &user.Password)
	if err != nil {
		return false, err
	}
	if userIn.Username == user.Username && userIn.Password == user.Password {
		return true, nil
	} else {
		return false, err
	}
}

func InitUserRepoImpl(db *sql.DB) UserRepository {
	return &UserRepoImpl{db: db}
}
