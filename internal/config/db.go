package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitializeDB() {
	connStr := `user=user password=roadguard dbname=roadguard sslmode=disable`
	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error testing database connection:", err)
	}

	fmt.Println("Database connection established")
}
