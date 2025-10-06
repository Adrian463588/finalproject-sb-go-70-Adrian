package main

import (
	"log"
	"os"

	"go-notes-api/database"
	"go-notes-api/handler"
	"go-notes-api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Muat environment dari file .env (opsional)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Tidak menemukan file .env (ini normal di production)")
	}

	// Hubungkan ke database
	database.Connect()

	// Set mode Gin
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)

	// Buat router baru
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// ====== ROUTES ======

	// Public routes
	public := router.Group("/api/users")
	{
		public.POST("/register", handler.RegisterUser)
		public.POST("/login", handler.LoginUser)
	}

	// Protected routes (JWT required)
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users/profile", handler.GetUserProfile)

		notes := protected.Group("/notes")
		{
			notes.GET("", handler.GetAllNotes)
			notes.POST("", handler.CreateNote)
			notes.GET("/favorites", handler.GetFavoriteNotes)
			notes.GET("/:id", handler.GetNoteByID)
			notes.PUT("/:id", handler.UpdateNote)
			notes.DELETE("/:id", handler.DeleteNote)
			notes.PUT("/:id/favorite", handler.ToggleFavoriteNote)
		}
	}

	// Jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server berjalan di port :%s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Gagal menjalankan server: %v", err)
	}
}
