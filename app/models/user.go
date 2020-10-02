package models

import (
	"gorm.io/gorm"
)

// User struct.
type User struct {
	FullName string `json:"full_name"`
	Email    string
	Password string
	OfficeID uint   `json:"office_id"`
	Office   Office `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	gorm.Model
}
