package handler

import (
	"database/sql"
	"go-notes-api/auth"
	"go-notes-api/database"
	"go-notes-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser mendaftarkan pengguna baru
func RegisterUser(c *gin.Context) {
	var user models.User

	// Binding dan validasi input JSON ke struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hashing password dengan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// SQL untuk memasukkan user baru dan mengembalikan ID-nya
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"
	err = database.DB.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		// Menangani kemungkinan email atau username sudah ada
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user, email or username may already exist"})
		return
	}

	// Jangan kirim balik password hash dalam respons
	user.Password = ""

	c.JSON(http.StatusCreated, user)
}

// LoginUser mengautentikasi pengguna dan mengembalikan token JWT
func LoginUser(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Cari user berdasarkan email
	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	err := database.DB.QueryRow(query, input.Email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Bandingkan password yang di-hash dengan password dari input
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		// Jika password tidak cocok, bcrypt akan mengembalikan error
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Jika password cocok, buat token JWT
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetUserProfile mengambil profil user yang sedang login
func GetUserProfile(c *gin.Context) {
	// Ambil userID dari context yang sudah disisipkan oleh middleware
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var user models.User
	query := "SELECT id, username, email FROM users WHERE id = $1"

	err := database.DB.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user profile"})
		return
	}

	c.JSON(http.StatusOK, user)
}