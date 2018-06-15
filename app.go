package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(username string, password string, dburl string) {
	connectionParams := fmt.Sprintf("user=%s password=%s dburl=%s", username, password, dburl)
	var err error

	a.DB, err = sql.Open("postgres", connectionParams)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/people", GetPeople).Methods("GET")
	a.Router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	a.Router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	a.Router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
}
