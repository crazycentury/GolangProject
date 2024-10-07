package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID        uuid.UUID
	TargetUserID  uuid.UUID
	Amount        int64
	Type          string // CREDIT/DEBIT
	Remarks       string
	CreatedAt     time.Time
}
