package main

import (
	"todo/postgres"
)

func main() {
	// pointer to postgres option
	DB := postgres.New(&pg.Options{
		User: 'postgres',
		Password: 'postgres',
		Database: 'postgres'
	})
}