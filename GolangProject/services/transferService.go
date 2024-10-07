package services

import (
	"GolangProject/db"
	"GolangProject/models"
	"time"

	"github.com/google/uuid"
)

type TransferRequest struct {
	FromUserID uuid.UUID
	ToUserID   uuid.UUID
	Amount     int64
	Remarks    string
}

var TransferQueue = make(chan TransferRequest, 100)

func StartTransferBackground() {
	for transfer := range TransferQueue {
		ExecuteTransfer(transfer)
	}
}

func ExecuteTransfer(request TransferRequest) {
	var fromUser, toUser models.User
	db.DB.First(&fromUser, "user_id = ?", request.FromUserID)
	db.DB.First(&toUser, "user_id = ?", request.ToUserID)

	if fromUser.Balance >= request.Amount {
		fromUser.Balance -= request.Amount
		toUser.Balance += request.Amount

		db.DB.Save(&fromUser)
		db.DB.Save(&toUser)

		transaction := models.Transaction{
			TransactionID: uuid.New(),
			UserID:        fromUser.UserID,
			TargetUserID:  toUser.UserID,
			Amount:        request.Amount,
			Type:          "DEBIT",
			Remarks:       request.Remarks,
			CreatedAt:     time.Now(),
		}
		db.DB.Create(&transaction)
	}
}
