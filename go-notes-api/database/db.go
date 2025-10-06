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

	// Railway menyediakan DATABASE_URL secara otomatis
	if os.Getenv("DATABASE_URL") != "" {
		connStr = os.Getenv("DATABASE_URL")

		// Tambahkan sslmode=require jika belum ada di URL
		if !containsSSLMode(connStr) {
			connStr += "?sslmode=require"
		}
	} else {
		// Untuk koneksi lokal tanpa SSL
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
	return len(connStr) >= 12 && ( // minimal panjang URL
		contains(connStr, "sslmode=") ||
			contains(connStr, "?sslmode="))
}

func contains(s, substr string) bool {
	if len(s) < len(substr) {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
