package main

import (
	"fmt"
	"database/sql"
)

type user struct {
	ID int`json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (u *user) String() string {
	return fmt.Sprintf("id:%d, first_name:%s, last_name:%s, email:%s, phone_number:%s",
		u.ID, u.FirstName, u.LastName, u.Email, u.PhoneNumber)
}

func (u *user) getUser(db *sql.DB) error {
	queryString := "SELECT first_name, last_name, email, phone_number FROM users WHERE id = $1"
	return db.QueryRow(queryString, u.ID).Scan(&u.FirstName, &u.LastName, &u.Email, &u.PhoneNumber)
}