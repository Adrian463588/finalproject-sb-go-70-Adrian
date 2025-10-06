package handler

import (
	"database/sql"
	"go-notes-api/database"
	"go-notes-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateNote menambahkan note baru ke database
func CreateNote(c *gin.Context) {
	var newNote models.Note
	if err := c.ShouldBindJSON(&newNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil userID dari context yang di-set oleh middleware
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Simpan note ke database dengan user_id yang sesuai
	sqlStatement := `INSERT INTO notes (user_id, title, content) VALUES ($1, $2, $3) RETURNING id`
	newNote.UserID = userID.(int)

	err := database.DB.QueryRow(sqlStatement, newNote.UserID, newNote.Title, newNote.Content).Scan(&newNote.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusCreated, newNote)
}

// GetAllNotes mengambil semua notes milik user yang sedang login
func GetAllNotes(c *gin.Context) {
	userID, _ := c.Get("userID")

	rows, err := database.DB.Query("SELECT id, user_id, title, content, is_favorite FROM notes WHERE user_id = $1", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notes"})
		return
	}
	defer rows.Close()

	notes := make([]models.Note, 0)
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.ID, &note.UserID, &note.Title, &note.Content, &note.IsFavorite); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan note data"})
			return
		}
		notes = append(notes, note)
	}

	c.JSON(http.StatusOK, notes)
}

// GetNoteByID mengambil satu note spesifik milik user
func GetNoteByID(c *gin.Context) {
	noteID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	userID, _ := c.Get("userID")
	var note models.Note

	sqlStatement := `SELECT id, user_id, title, content, is_favorite FROM notes WHERE id = $1 AND user_id = $2`
	err = database.DB.QueryRow(sqlStatement, noteID, userID).Scan(&note.ID, &note.UserID, &note.Title, &note.Content, &note.IsFavorite)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve note"})
		return
	}

	c.JSON(http.StatusOK, note)
}

// UpdateNote memperbarui note yang sudah ada
func UpdateNote(c *gin.Context) {
	noteID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	var updatedNote models.Note
	if err := c.ShouldBindJSON(&updatedNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")

	sqlStatement := `UPDATE notes SET title = $1, content = $2 WHERE id = $3 AND user_id = $4`
	result, err := database.DB.Exec(sqlStatement, updatedNote.Title, updatedNote.Content, noteID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to edit this note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note updated successfully"})
}

// DeleteNote menghapus note milik user
func DeleteNote(c *gin.Context) {
	noteID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	userID, _ := c.Get("userID")

	sqlStatement := `DELETE FROM notes WHERE id = $1 AND user_id = $2`
	result, err := database.DB.Exec(sqlStatement, noteID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Note not found or you don't have permission to delete it"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}

// GetFavoriteNotes mengambil semua notes yang ditandai sebagai favorit
func GetFavoriteNotes(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	query := "SELECT id, user_id, title, content, is_favorite FROM notes WHERE user_id = $1 AND is_favorite = TRUE ORDER BY updated_at DESC"
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve favorite notes"})
		return
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.ID, &note.UserID, &note.Title, &note.Content, &note.IsFavorite); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan note data"})
			return
		}
		notes = append(notes, note)
	}

	c.JSON(http.StatusOK, notes)
}

// ToggleFavoriteNote mengubah status is_favorite pada sebuah note
func ToggleFavoriteNote(c *gin.Context) {
	noteID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	userID, _ := c.Get("userID")

	query := "UPDATE notes SET is_favorite = NOT is_favorite, updated_at = NOW() WHERE id = $1 AND user_id = $2"
	result, err := database.DB.Exec(query, noteID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update favorite status"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check affected rows"})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Note not found or you don't have permission to change it"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note favorite status toggled"})
}
