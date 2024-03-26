package models

import (
	"log"

	"example.com/rest-api/db"
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

	result, err:= stmt.Exec(u.Name, u.Email, u.Password)
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
		log.Printf("Error retrieving event: %v", err)
		return err
	}
	users = append(users, *u)
	return nil
}