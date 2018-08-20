package main

import "fmt"

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