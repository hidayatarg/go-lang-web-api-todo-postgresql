package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo/handler"
	"todo/postgres"

	"github.com/go-pg/pg/v9"
)

func main() {
	// pointer to postgres option
	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	})

	defer DB.Close()

	r := handler.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatalf("cannot start server %v", err)
	}

}
