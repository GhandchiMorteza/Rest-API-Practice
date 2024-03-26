package models

import (
	"log"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Name     string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

var users []User

func (u *User) Save() error {
	query := `INSERT INTO users(name, email, password)
	Values (?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err:= stmt.Exec(u.Name, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}

func (u *User) SaveAndUpdateUsers() error {
	err := u.Save()
	if err != nil {
		log.Printf("Error saving user: %v", err)
		return err
	}
	users = append(users, *u)
	return nil
}