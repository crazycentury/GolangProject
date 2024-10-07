package controllers

import (
	"GolangProject/db"
	"GolangProject/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("secret_key")

// Register Controller
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Cek apakah nomor telepon sudah terdaftar
	var existingUser models.User
	db.DB.Where("phone_number = ?", user.PhoneNumber).First(&existingUser)
	if existingUser.UserID != uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Phone Number already registered"})
		return
	}

	hashedPIN, _ := bcrypt.GenerateFromPassword([]byte(user.PIN), 10)
	user.PIN = string(hashedPIN)
	user.UserID = uuid.New()
	db.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": user,
	})
}

// Login Controller
func Login(c *gin.Context) {
	var user models.User
	var request struct {
		PhoneNumber string
		PIN         string
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Cek apakah nomor telepon dan PIN cocok
	db.DB.Where("phone_number = ?", request.PhoneNumber).First(&user)
	if user.UserID == uuid.Nil || bcrypt.CompareHashAndPassword([]byte(user.PIN), []byte(request.PIN)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Phone Number and PIN doesn’t match."})
		return
	}

	// Generate JWT Token
	accessToken := generateToken(user.UserID.String())
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": map[string]string{
			"access_token": accessToken,
		},
	})
}

// Fungsi untuk generate JWT
func generateToken(userID string) string {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}
