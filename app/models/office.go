package models

import (
	"orm/app/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Office struct.
type Office struct {
	Name    string
	Address string
	gorm.Model
}

// MakeResponse of Office.
func (office Office) MakeResponse() gin.H {

	response := gin.H{
		"id":         office.ID,
		"name":       office.Name,
		"address":    office.Address,
		"created_at": office.CreatedAt.Format(helper.YMDHIS),
		"updated_at": office.UpdatedAt.Format(helper.YMDHIS),
	}

	return response
}

// MakeResponseWithUser of Office.
func (office Office) MakeResponseWithUser() gin.H {
	var (
		user User
	)
	response := gin.H{
		"id":         office.ID,
		"name":       office.Name,
		"address":    office.Address,
		"created_at": office.CreatedAt.Format(helper.YMDHIS),
		"updated_at": office.UpdatedAt.Format(helper.YMDHIS),
		"users":      user.GetByOfficeID(office.ID, false),
	}

	return response
}

// MakeResponseWithToDo of Office.
func (office Office) MakeResponseWithToDo() gin.H {
	var (
		user User
	)
	response := gin.H{
		"id":         office.ID,
		"name":       office.Name,
		"address":    office.Address,
		"created_at": office.CreatedAt.Format(helper.YMDHIS),
		"updated_at": office.UpdatedAt.Format(helper.YMDHIS),
		"users":      user.GetByOfficeID(office.ID, true),
	}

	return response
}
