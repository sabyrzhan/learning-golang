package models

import (
	"rest-api/db"
	"rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?,?)"
	prepare, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer prepare.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	exec, err := prepare.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = id

	return nil
}
