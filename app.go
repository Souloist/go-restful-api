package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App contains reference to db interface and multiplexer
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize initalizes database connection and defines routes on multiplexer
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

// Run serves the app
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/people", GetPersonsHandler).Methods("GET")
	a.Router.HandleFunc("/people/{id}", GetPersonHandler).Methods("GET")
	a.Router.HandleFunc("/people/{id}", CreatePersonHandler).Methods("POST")
	a.Router.HandleFunc("/people/{id}", UpdatePersonHandler).Methods("PUT")
	a.Router.HandleFunc("/people/{id}", DeletePersonHandler).Methods("DELETE")
}
