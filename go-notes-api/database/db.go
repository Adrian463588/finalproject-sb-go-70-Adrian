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
	var connStr string
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL != "" {
		// Tambahkan sslmode=require jika belum ada
		if !strings.Contains(dbURL, "sslmode=") {
			dbURL += "?sslmode=require"
		}
		connStr = dbURL
		fmt.Println("📦 Menggunakan DATABASE_URL dari environment Railway")
	} else {
		fmt.Println("⚠️ DATABASE_URL tidak ditemukan, menggunakan konfigurasi lokal (.env)")
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
