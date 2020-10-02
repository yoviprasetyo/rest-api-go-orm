package controllers

import (
	"fmt"
	"net/http"
	"orm/app/models"

	"github.com/gin-gonic/gin"
)

// CreateToDo controller.
func (strDB *StrDB) CreateToDo(c *gin.Context) {
	var (
		todo   models.ToDo
		result gin.H
	)

	err := c.Bind(&todo)
	if err != nil {
		fmt.Println("tidak ada data")
	}

	strDB.DB.Create(&todo)
	result = responseAPI(todo, 1)

	c.JSON(http.StatusOK, result)

}

// DeleteToDo controller.
func (strDB *StrDB) DeleteToDo(c *gin.Context) {
	var (
		todo       models.ToDo
		result     gin.H
		resultCode int
	)

	id := c.Param("id")
	err := strDB.DB.First(&todo, id).Error
	resultCode = http.StatusOK
	result = gin.H{
		"ok":      true,
		"message": "Data deleted",
	}
	if err != nil {
		result = gin.H{
			"ok":      false,
			"message": "Data not found",
		}
		resultCode = http.StatusNotFound
	}

	errDelete := strDB.DB.Delete(&todo, id).Error
	if errDelete != nil {
		result = gin.H{
			"ok":      false,
			"message": errDelete.Error,
		}
		resultCode = http.StatusInternalServerError
	}

	c.JSON(resultCode, result)
}

// UpdateToDo controller.
func (strDB StrDB) UpdateToDo(c *gin.Context) {

	var (
		todo       models.ToDo
		newToDo    models.ToDo
		result     gin.H
		resultCode = http.StatusOK
	)

	result = gin.H{
		"ok":      true,
		"message": "Success update",
	}

	id := c.Param("id")
	err := c.Bind(&newToDo)
	err = strDB.DB.First(&todo, id).Error
	if err != nil {
		result = gin.H{
			"ok":      false,
			"message": "Data not found",
		}
		resultCode = http.StatusNotFound
	}

	if resultCode != http.StatusNotFound {
		errUpdate := strDB.DB.Model(&todo).Updates(newToDo).Error
		if errUpdate != nil {
			result = gin.H{
				"ok":      false,
				"message": errUpdate.Error,
			}
			resultCode = http.StatusInternalServerError
		}
	}

	c.JSON(resultCode, result)
}

// GetToDo controllers.
func (strDB *StrDB) GetToDo(c *gin.Context) {
	var (
		todo   []models.ToDo
		result gin.H
	)

	strDB.DB.Find(&todo)

	if len(todo) <= 0 {
		arrayNil := []models.ToDo{}
		result = gin.H{
			"result": arrayNil,
			"count":  0,
		}
	}

	if len(todo) > 0 {
		result = responseAPI(todo, len(todo))
	}

	c.JSON(http.StatusOK, result)
}

// GetOneToDo controllers.
func (strDB *StrDB) GetOneToDo(c *gin.Context) {
	var (
		todo       models.ToDo
		resultCode = http.StatusOK
		result     gin.H
	)

	id := c.Param("id")
	err := strDB.DB.First(&todo, id).Error
	if err != nil {
		resultCode = http.StatusNotFound
		result = gin.H{
			"ok":      false,
			"message": "Data not found",
		}
	}

	if resultCode == http.StatusOK {
		result = responseAPI(todo, 1)
	}

	c.JSON(resultCode, result)

}

// GetSearchToDo controllers.
func (strDB *StrDB) GetSearchToDo(c *gin.Context) {
	var (
		todo       []models.ToDo
		resultCode = http.StatusOK
		result     gin.H
	)

	search := c.Query("search")
	strDBQuery := strDB.DB

	if search == "" {
		resultCode = http.StatusBadRequest
		result = gin.H{
			"ok":      false,
			"message": "Search keyword is required",
		}
	}

	if search != "" {
		err := strDBQuery.Where("name LIKE ?", "%"+search+"%").Or("description LIKE ?", "%"+search+"%").Find(&todo).Error
		if err != nil {
			resultCode = http.StatusNotFound
			result = gin.H{
				"ok":      false,
				"message": err.Error,
			}
		}
	}

	if resultCode == http.StatusOK {
		result = responseAPI(todo, len(todo))
	}

	c.JSON(resultCode, result)

}
