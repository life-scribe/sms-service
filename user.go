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

func (u *user) createUser(db *sql.DB) error {
	queryString := "INSERT INTO users (first_name, last_name, email, phone_number) VALUES ($1, $2, $3, $4)"
	if err := db.QueryRow(queryString, u.FirstName, u.LastName, u.Email, u.PhoneNumber).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}

func (u *user) deleteUser(db *sql.DB) error {
	queryString := "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(queryString, u.ID)

	return err
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	queryString := "SELECT * FROM users LIMIT $1 OFFSET $2"
	users := []user{}

	if rows, err := db.Query(queryString, count, start); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		for rows.Next() {
			var u user
			if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.PhoneNumber); err != nil {
				return nil, err
			}
			users = append(users, u)
		}
		return users, nil
	}
}

func (u * user) updateUser(db *sql.DB) error {
	queryString := "UDPATE users SET first_name = $1, last_name = $2, email = $3, phone_number = $4 WHERE id = $5"
	_, err := db.Exec(queryString, u.FirstName, u.LastName, u.Email, u.PhoneNumber, u.ID)

	return err
}