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
	// Load .env hanya jika file ada (di lokal)
	_ = godotenv.Load()

	// Hubungkan ke database
	if err := database.Connect(); err != nil {
		log.Fatalf("❌ Gagal koneksi ke database: %v", err)
	}

	// Gunakan mode release jika tidak diatur
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server berjalan di port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Gagal menjalankan server: %v", err)
	}
}
