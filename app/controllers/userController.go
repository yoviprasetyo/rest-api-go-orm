package controllers

/*
import (
	"fmt"
	"net/http"
	"orm/app/helper"
	"orm/app/models"

	"github.com/gin-gonic/gin"
)

// CreateUser controller.
func (controller *Controller) CreateUser(c *gin.Context) {
	var (
		user   models.User
		result gin.H
	)

	err := c.Bind(&user)
	if err != nil {
		fmt.Println("tidak ada data")
	}

	hashed, err := helper.HashPassword(user.Password)
	if err != nil {
		result = gin.H{
			"status":  "error",
			"message": err.Error,
		}
		c.JSON(http.StatusInternalServerError, result)
		c.Abort()
	}

	user.Password = hashed

	controller.DB.DB.Create(&user)
	controller.DB.DB.Joins("Office").First(&user, user.ID)

	result = responseAPI(user, "")

	c.JSON(http.StatusOK, result)

}

// DeleteUser controller.
func (controller *Controller) DeleteUser(c *gin.Context) {
	var (
		user       models.User
		result     gin.H
		resultCode int
	)

	id := c.Param("id")
	err := controller.DB.DB.First(&user, id).Error
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

	errDelete := controller.DB.DB.Delete(&user, id).Error
	if errDelete != nil {
		result = gin.H{
			"ok":      false,
			"message": errDelete.Error,
		}
		resultCode = http.StatusInternalServerError
	}

	c.JSON(resultCode, result)
}

// UpdateUser controller.
func (controller *Controller) UpdateUser(c *gin.Context) {

	var (
		user       models.User
		newUser    models.User
		result     gin.H
		resultCode = http.StatusOK
	)

	result = gin.H{
		"ok":      true,
		"message": "Success update",
	}

	id := c.Param("id")
	err := c.Bind(&newUser)
	err = controller.DB.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"ok":      false,
			"message": "Data not found",
		}
		resultCode = http.StatusNotFound
	}

	if resultCode != http.StatusNotFound {
		errUpdate := controller.DB.DB.Model(&user).Updates(newUser).Error
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

// GetUser controllers.
func (controller *Controller) GetUser(c *gin.Context) {
	var (
		user   []models.User
		result gin.H
	)

	controller.DB.DB.Find(&user)

	if len(user) <= 0 {
		arrayNil := []models.User{}
		result = gin.H{
			"result": arrayNil,
			"count":  0,
		}
	}

	if len(user) > 0 {
		result = responseAPI(user, "")
	}

	c.JSON(http.StatusOK, result)
}

// GetOneUser controllers.
func (controller *Controller) GetOneUser(c *gin.Context) {
	var (
		user       models.User
		resultCode = http.StatusOK
		result     gin.H
	)

	id := c.Param("id")
	err := controller.DB.DB.First(&user, id).Error
	if err != nil {
		resultCode = http.StatusNotFound
		result = gin.H{
			"ok":      false,
			"message": "Data not found",
		}
	}

	if resultCode == http.StatusOK {
		result = responseAPI(user, "")
	}

	c.JSON(resultCode, result)

}

// GetSearchUser controllers.
func (controller *Controller) GetSearchUser(c *gin.Context) {
	var (
		user       []models.User
		resultCode = http.StatusOK
		result     gin.H
	)

	search := c.Query("search")
	strDBQuery := controller.DB.DB

	if search == "" {
		resultCode = http.StatusBadRequest
		result = gin.H{
			"ok":      false,
			"message": "Search keyword is required",
		}
	}

	if search != "" {
		err := strDBQuery.Where("full_name LIKE ?", "%"+search+"%").Or("email LIKE ?", "%"+search+"%").Find(&user).Error
		if err != nil {
			resultCode = http.StatusNotFound
			result = gin.H{
				"ok":      false,
				"message": err.Error,
			}
		}
	}

	if resultCode == http.StatusOK {
		result = responseAPI(user, "")
	}

	c.JSON(resultCode, result)

}
*/
