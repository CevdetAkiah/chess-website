package main

import (
	"fmt"
	postgres "go-projects/chess/database"

	"log"
)

func main() {
	err := postgres.Psql.Ping()
	if err != nil {
		err = fmt.Errorf("Cannot connect to database with error: %w", err)
		log.Fatalln(err)
	}

	fmt.Println("connected to database website")
}
