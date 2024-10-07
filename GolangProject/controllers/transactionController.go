package controllers

import (
	"GolangProject/db"
	"GolangProject/models"
	"GolangProject/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TopUp(c *gin.Context) {
	var request struct {
		Amount int64
	}
	userID := c.GetString("user_id")

	// Assume user is authenticated from JWT
	var user models.User
	db.DB.First(&user, "user_id = ?", userID)

	balanceBefore := user.Balance
	user.Balance += request.Amount
	db.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"top_up_id":      uuid.New(),
			"amount_top_up":  request.Amount,
			"balance_before": balanceBefore,
			"balance_after":  user.Balance,
		},
	})
}

func Transfer(c *gin.Context) {
	var request struct {
		TargetUserID string
		Amount       int64
		Remarks      string
	}
	userID := c.GetString("user_id")

	transferRequest := services.TransferRequest{
		FromUserID: uuid.MustParse(userID),
		ToUserID:   uuid.MustParse(request.TargetUserID),
		Amount:     request.Amount,
		Remarks:    request.Remarks,
	}

	services.TransferQueue <- transferRequest
	c.JSON(http.StatusOK, gin.H{"status": "SUCCESS", "message": "Transfer in progress"})
}
