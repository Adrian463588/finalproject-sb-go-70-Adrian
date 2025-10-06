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

func Connect() {
	dbURL := os.Getenv("DATABASE_URL")
	var connStr string

	if dbURL != "" {
		// Pastikan sslmode diset (Railway butuh sslmode=require)
		if !strings.Contains(dbURL, "sslmode=") {
			if strings.Contains(dbURL, "?") {
				dbURL += "&sslmode=require"
			} else {
				dbURL += "?sslmode=require"
			}
		}
		connStr = dbURL
		log.Println("📦 Menggunakan DATABASE_URL dari environment")
	} else {
		// Fallback ke konfigurasi lokal (biasanya digunakan saat development)
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
		log.Fatalf("❌ Gagal membuka koneksi database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("❌ Database tidak bisa dijangkau: %v", err)
	}

	DB = db
	log.Println("✅ Berhasil terhubung ke database!")
}
