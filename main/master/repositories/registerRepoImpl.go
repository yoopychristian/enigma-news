package repositories

import (
	"database/sql"
	"enigma-news/main/master/constantaQuery"
	"enigma-news/main/master/models"
	"log"
)

type RegisterRepoImpl struct {
	db *sql.DB
}

func (s RegisterRepoImpl) GetAllRegisters() ([]*models.Register, error) {
	dataRegister := []*models.Register{}
	query := constantaQuery.GETALLREGISTER
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		register := models.Register{}
		var err = data.Scan(&register.Id, &register.Email, &register.Fullname, &register.Phonenumber, &register.Username, &register.Password)
		if err != nil {
			return nil, err
		}
		dataRegister = append(dataRegister, &register)
	}
	return dataRegister, nil
}

func (s RegisterRepoImpl) GetRegistersID(Username string) (*models.Register, error) {
	register := new(models.Register)
	query := constantaQuery.GETREGISTERBYUNAME
	if err := s.db.QueryRow(query, Username).Scan(&register.Id, &register.Email, &register.Fullname, &register.Phonenumber, &register.Username, &register.Password); err != nil {
		return nil, err
	}
	return register, nil
}

func (s RegisterRepoImpl) CreateRegisters(register models.Register) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := constantaQuery.GETADDREGISTER

	stmt, err := s.db.Prepare(query)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(register.Id, register.Email, register.Phonenumber, register.Username, register.Password); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return nil
}

func (s RegisterRepoImpl) UpdateRegisters(register models.Register) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := constantaQuery.GETUPDATEREGISTER

	stmt, err := s.db.Prepare(query)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(register.Email, register.Phonenumber, register.Username, register.Password, register.Id); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return nil
}

func (s RegisterRepoImpl) DeleteRegisters(Username string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := constantaQuery.GETUPDATEREGISTER

	stmt, err := s.db.Prepare(query)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(Username); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return nil
}

func InitRegisterRepoImpl(db *sql.DB) RegisterRepository {
	return &RegisterRepoImpl{db}
}
