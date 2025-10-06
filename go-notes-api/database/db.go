// database/connect.go (atau file database.go kamu yang sekarang)
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
		// Pastikan sslmode ada
		if !strings.Contains(dbURL, "sslmode=") {
			// Jika URL sudah punya query params “?”, tambahkan “&”… tapi untuk sederhana:
			dbURL += "?sslmode=require"
		}
		connStr = dbURL
		fmt.Println("📦 Menggunakan DATABASE_URL dari environment")
	} else {
		fmt.Println("⚠️ DATABASE_URL tidak ditemukan, fallback ke konfigurasi lokal")
		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		)
	}

	// Debug print koneksi string (hati-hati, ini bisa mengekspos password, gunakan sementara saja)
	fmt.Println("ConnString:", connStr)

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
