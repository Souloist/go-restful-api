package main

import "database/sql"

// GetPerson returns person row given id
func GetPerson(db *sql.DB) {}

// GetPersons returns all person rows or subset if filter is provided
func GetPersons(db *sql.DB) []Person {
	return []Person{}
}

// CreatePerson writes person entry to db
func CreatePerson(db *sql.DB) {}

// UpdatePerson updates person row in db
func UpdatePerson(db *sql.DB) {}

// DeletePerson deletes person row in db
func DeletePerson(db *sql.DB) {}
