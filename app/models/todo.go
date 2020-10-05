package models

import (
	"orm/app/helper"

	"github.com/gin-gonic/gin"
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

// MakeResponse of ToDo.
func (toDo ToDo) MakeResponse() gin.H {
	response := gin.H{
		"id":          toDo.ID,
		"name":        toDo.Name,
		"description": toDo.Description,
		"created_at":  toDo.CreatedAt.Format(helper.YMDHIS),
		"updated_at":  toDo.UpdatedAt.Format(helper.YMDHIS),
	}

	return response
}

// MakeResponseWithUser of ToDo.
func (toDo ToDo) MakeResponseWithUser() gin.H {
	DB.Preload("User").First(&toDo, toDo.ID)
	response := gin.H{
		"id":          toDo.ID,
		"name":        toDo.Name,
		"description": toDo.Description,
		"created_at":  toDo.CreatedAt.Format(helper.YMDHIS),
		"updated_at":  toDo.UpdatedAt.Format(helper.YMDHIS),
		"user":        toDo.User.MakeResponse(),
	}

	return response
}

// MakeResponseWithOffice of ToDo.
func (toDo ToDo) MakeResponseWithOffice() gin.H {
	DB.Preload("User").First(&toDo, toDo.ID)
	response := gin.H{
		"id":          toDo.ID,
		"name":        toDo.Name,
		"description": toDo.Description,
		"created_at":  toDo.CreatedAt.Format(helper.YMDHIS),
		"updated_at":  toDo.UpdatedAt.Format(helper.YMDHIS),
		"user":        toDo.User.MakeResponseWithOffice(),
	}

	return response
}

// GetByUserID method.
func (toDo ToDo) GetByUserID(id uint) []gin.H {
	var (
		g     []gin.H
		toDos []ToDo
	)
	DB.Where("user_id = ?", id).Find(&toDos)
	for _, toDo := range toDos {
		g = append(g, toDo.MakeResponse())
	}
	return g
}
