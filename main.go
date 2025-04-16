package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go_with_mysql/database"
	"log"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
}
