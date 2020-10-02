package models

import (
	"gorm.io/gorm"
)

// Office struct.
type Office struct {
	Name    string
	Address string
	gorm.Model
}

// OfficeResponse struct.
type OfficeResponse struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
