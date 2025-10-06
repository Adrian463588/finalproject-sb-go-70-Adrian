// file: auth/jwt.go
package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Ganti ini dengan secret key yang lebih aman dan simpan di .env
var jwtKey = []byte("your_super_secret_key")

// Claims adalah data yang kita simpan di dalam token
type JWTClaim struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken membuat token JWT baru untuk user ID tertentu
func GenerateToken(userID int) (string, error) {
	// Token berlaku selama 24 jam
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateToken memverifikasi token dan mengembalikan claims jika valid
func ValidateToken(signedToken string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}