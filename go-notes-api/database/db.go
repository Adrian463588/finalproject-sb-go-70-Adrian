package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	dbURL := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	var connStr string

	if dbURL != "" {
		// Pastikan format sesuai PostgreSQL
		if strings.HasPrefix(dbURL, "postgres://") {
			dbURL = strings.Replace(dbURL, "postgres://", "postgresql://", 1)
		}

		// Tambah sslmode jika belum ada
		if !strings.Contains(dbURL, "sslmode=") {
			if strings.Contains(dbURL, "?") {
				dbURL += "&sslmode=require"
			} else {
				dbURL += "?sslmode=require"
			}
		}

		connStr = dbURL
		log.Println("📦 Menggunakan DATABASE_URL dari Railway environment")
	} else {
		// Fallback lokal (dev only)
		log.Println("⚠️  DATABASE_URL tidak ditemukan, menggunakan konfigurasi lokal")
		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("gagal membuka koneksi database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("database tidak bisa dijangkau: %v", err)
	}

	DB = db
	log.Println("✅ Database berhasil terhubung!")
	return nil
}
