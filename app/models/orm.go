package models

import "gorm.io/gorm"

// Additional struct.
type Additional struct {
	Key     string
	Content interface{}
}

// DB var.
var DB *gorm.DB
