package models

import (
	"orm/app/helper"

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

// MakeResponse of Office.
func (office Office) MakeResponse() OfficeResponse {
	return OfficeResponse{
		Name:      office.Name,
		Address:   office.Address,
		CreatedAt: office.CreatedAt.Format(helper.YMDHIS),
		UpdatedAt: office.UpdatedAt.Format(helper.YMDHIS),
	}
}
