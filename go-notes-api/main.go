// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"go-notes-api/database"
	"go-notes-api/handler"
	"go-notes-api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Muat variabel dari .env (jika ada) — tapi ini hanya berlaku lokal / jika ada file .env di container
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Error loading .env file:", err)
	}

	// Debug print variabel penting (sementara)
	fmt.Println("---- ENV VARS ----")
	fmt.Println("DATABASE_URL =", os.Getenv("DATABASE_URL"))
	fmt.Println("DB_HOST =", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT =", os.Getenv("DB_PORT"))
	fmt.Println("PORT =", os.Getenv("PORT"))
	fmt.Println("------------------")

	// Hubungkan ke database
	database.Connect()

	// Set mode Gin — harus dilakukan *sebelum* membuat router
	gin.SetMode(gin.ReleaseMode) // agar tidak dalam debug mode di produksi
	// Alternatif: bisa set GIN_MODE=release di env Railway juga

	router := gin.New()
	// Tambahkan middleware bawaan: logger & recovery
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Routes
	public := router.Group("/api/users")
	{
		public.POST("/register", handler.RegisterUser)
		public.POST("/login", handler.LoginUser)
	}

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users/profile", handler.GetUserProfile)

		protected.GET("/notes", handler.GetAllNotes)
		protected.POST("/notes", handler.CreateNote)
		protected.GET("/notes/favorites", handler.GetFavoriteNotes)
		protected.GET("/notes/:id", handler.GetNoteByID)
		protected.PUT("/notes/:id", handler.UpdateNote)
		protected.DELETE("/notes/:id", handler.DeleteNote)
		protected.PUT("/notes/:id/favorite", handler.ToggleFavoriteNote)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server akan berjalan di port :%s\n", port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("❌ Gagal menjalankan server:", err)
	}
}
