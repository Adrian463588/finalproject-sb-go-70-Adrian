// file: main.go
package main

import (
	"go-notes-api/database"
	"go-notes-api/handler"
	"go-notes-api/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Muat variabel dari .env di awal aplikasi
	err := godotenv.Load() // <-- TAMBAHKAN BARIS INI
	if err != nil {
		log.Println("Error loading .env file")
	}
	
	database.Connect()

	router := gin.Default()

	// Grup untuk route publik (tidak perlu token)
	public := router.Group("/api/users")
	{
		public.POST("/register", handler.RegisterUser)
		public.POST("/login", handler.LoginUser)
	}

	// Grup untuk route yang dilindungi JWT
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
	log.Fatal(router.Run(":" + port))
}