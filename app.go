package main

import (
	"github.com/gorilla/mux"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"encoding/json"
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

	a.Router = mux.NewRouter()
	a.initRoutes()
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
	a.Router.HandleFunc("/user", a.createUser).Methods("POST")
	a.Router.HandleFunc("/user/{id:[0-9]+}", a.getUser).Methods("GET")
	a.Router.HandleFunc("/user/{id:[0-9]+}", a.updateUser).Methods("PUT")
	a.Router.HandleFunc("/user/{id:[0-9]+}", a.deleteUser).Methods("DELETE")
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
	users := []user{}
	rows := new(sql.Rows)
	var err error
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
	respondWithJSON(w, http.StatusOK, users)
}

func (a *App) getUser(w http.ResponseWriter, r *http.Request) {

}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {

}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {

}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}