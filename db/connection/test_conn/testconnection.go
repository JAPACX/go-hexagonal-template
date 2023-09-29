package main

import (
	"go-gqlgen/db/connection/dbconn"
	"log"
)

func main() {

	db := dbconn.GetDB() //
	err := db.Ping()
	if err != nil {
		log.Fatal("Error:", err)
	} else {
		log.Println("Successfully connected to the database!")
	}
}
