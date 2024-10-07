package models

import "github.com/google/uuid"

type User struct {
	UserID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName   string
	LastName    string
	PhoneNumber string `gorm:"unique"`
	Address     string
	PIN         string
	Balance     int64
}
