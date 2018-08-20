package main

import (
	"github.com/gorilla/mux"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Router *mux.Router
	DB *sql.DB
}

func (a *App) Init(username, password, db string) {
	connString := fmt.Sprintf("%s:%s@/%s?charset=utf8", username, password, db)
	var err error
	if a.DB, err = sql.Open("mysql", connString); err != nil {
		log.Fatal(err)
	}

	users := []user{}
	rows := new(sql.Rows)
	if rows, err = a.DB.Query("SELECT * FROM users"); err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.PhoneNumber); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	for _, u := range users {
		fmt.Println(u.String())
	}
}