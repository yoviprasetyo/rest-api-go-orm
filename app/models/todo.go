package models

import (
	"gorm.io/gorm"
)

// ToDo struct.
type ToDo struct {
	Name        string
	Description string
	UserID      uint `json:"user_id"`
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	gorm.Model
}
