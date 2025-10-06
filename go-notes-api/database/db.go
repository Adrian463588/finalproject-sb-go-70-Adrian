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
		// Railway membutuhkan SSL koneksi aktif
		connStr = os.Getenv("DATABASE_URL")

		// Tambahkan sslmode=require jika belum ada di URL
		if !containsSSLMode(connStr) {
			connStr += "?sslmode=require"
		}
	} else {
		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		)
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Gagal membuka koneksi ke database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Database tidak dapat dijangkau:", err)
	}

	fmt.Println("✅ Berhasil terhubung ke database!")
}

func containsSSLMode(connStr string) bool {
	return len(connStr) >= 12 && ( // minimal panjang url
		contains(connStr, "sslmode=") ||
			contains(connStr, "?sslmode="))
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && ( // simple string check
		len(s) >= len(substr) && (s == substr || (len(s) > len(substr) && (s[:len(substr)] == substr || contains(s[1:], substr)))))
}
