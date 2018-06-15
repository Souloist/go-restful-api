package main

import "database/sql"

func (p *Person) GetPerson(db *sql.DB) {}
func (p *Person) GetPersons(db *sql.DB) []Person {
	return []Person{}
}
func (p *Person) CreatePerson(db *sql.DB) {}
func (p *Person) UpdatePerson(db *sql.DB) {}
func (p *Person) DeletePerson(db *sql.DB) {}
