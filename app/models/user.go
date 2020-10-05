package models

import (
	"orm/app/helper"

	"github.com/gin-gonic/gin"
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

// MakeResponse of User.
func (user User) MakeResponse() gin.H {

	response := gin.H{
		"id":         user.ID,
		"full_name":  user.FullName,
		"email":      user.Email,
		"created_at": user.CreatedAt.Format(helper.YMDHIS),
		"updated_at": user.UpdatedAt.Format(helper.YMDHIS),
	}

	return response
}

// MakeResponseWithOffice of User.
func (user User) MakeResponseWithOffice() gin.H {
	DB.Preload("Office").First(&user, user.ID)
	response := gin.H{
		"id":         user.ID,
		"full_name":  user.FullName,
		"email":      user.Email,
		"created_at": user.CreatedAt.Format(helper.YMDHIS),
		"updated_at": user.UpdatedAt.Format(helper.YMDHIS),
		"office":     user.Office.MakeResponse(),
	}

	return response
}

// MakeResponseWithToDo of User.
func (user User) MakeResponseWithToDo() gin.H {
	var (
		toDo ToDo
	)
	response := gin.H{
		"id":         user.ID,
		"full_name":  user.FullName,
		"email":      user.Email,
		"created_at": user.CreatedAt.Format(helper.YMDHIS),
		"updated_at": user.UpdatedAt.Format(helper.YMDHIS),
		"to_dos":     toDo.GetByUserID(user.ID),
	}

	return response
}

// GetByOfficeID method.
func (user User) GetByOfficeID(id uint, withToDo bool) []gin.H {
	var (
		g     []gin.H
		users []User
	)
	DB.Where("office_id = ?", id).Find(&users)
	for _, user := range users {
		if withToDo {
			g = append(g, user.MakeResponseWithToDo())
		}
		if !withToDo {
			g = append(g, user.MakeResponse())
		}

	}
	return g
}
