// file: database/db.go
package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var connStr string
	if os.Getenv("DATABASE_URL") != "" {
		connStr = os.Getenv("DATABASE_URL")
	} else {
		connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil { log.Fatal("Gagal terhubung ke database:", err) }

	err = DB.Ping()
	if err != nil { log.Fatal("Database tidak dapat dijangkau:", err) }

	fmt.Println("Berhasil terhubung ke database!")
}