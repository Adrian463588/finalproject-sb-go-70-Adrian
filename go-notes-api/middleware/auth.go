// file: middleware/auth.go
package middleware

import (
	"go-notes-api/auth" // Pastikan package auth diimpor
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort() // Hentikan request
			return
		}

		// 2. Pisahkan "Bearer" dengan token-nya
		// Formatnya adalah "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 3. Panggil fungsi ValidateToken dari package auth
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// 4. Jika token valid, sisipkan userID ke dalam context
		// agar bisa digunakan oleh handler selanjutnya
		c.Set("userID", claims.UserID)

		// 5. Lanjutkan ke handler berikutnya
		c.Next()
	}
}