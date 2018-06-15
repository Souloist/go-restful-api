package main

import "net/http"

// GetPersonsHandler defines GET handler for /people/ route
func GetPersonsHandler(w http.ResponseWriter, r *http.Request) {}

// GetPersonHandler defines GET handler for /people/<id> route
func GetPersonHandler(w http.ResponseWriter, r *http.Request) {}

// CreatePersonHandler defines POST handler for /people/<id> route
func CreatePersonHandler(w http.ResponseWriter, r *http.Request) {}

// UpdatePersonHandler defines PUT handler for /people/<id> route
func UpdatePersonHandler(w http.ResponseWriter, r *http.Request) {}

// DeletePersonHandler defines DELETE handler for /people/<id> route
func DeletePersonHandler(w http.ResponseWriter, r *http.Request) {}
