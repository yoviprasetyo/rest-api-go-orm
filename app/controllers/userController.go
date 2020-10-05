package controllers

import (
	"errors"
	"net/http"
	"orm/app/helper"
	"orm/app/models"

	"github.com/gin-gonic/gin"
)

// CreateUser controller.
func (controller *Controller) CreateUser(c *gin.Context) {
	var (
		userResponse []gin.H
		user         models.User
		response     = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	controller.UseDB()

	err := c.Bind(&user)
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.ErrorMessage = err.Error()
		responseAPI(response)
	}

	errValidation := validateUser(user)
	if errValidation != nil {
		response.StatusCode = http.StatusBadRequest
		response.ErrorMessage = errValidation.Error()
		responseAPI(response)
	}

	hashed, err := helper.HashPassword(user.Password)
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.ErrorMessage = err.Error()
		responseAPI(response)
	}

	user.Password = hashed

	models.DB.Create(&user)

	userResponse = append(userResponse, user.MakeResponseWithOffice())
	response.Data = userResponse
	responseAPI(response)
}

/*
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

// GetUserOffice controllers.
func (controller *Controller) GetUserOffice(c *gin.Context) {
	var (
		userResponse []gin.H
		users        []models.User
		response     = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	controller.UseDB()

	OfficeID := c.Param("id")

	models.DB.Where("office_id = ?", OfficeID).Find(&users)
	for i := 0; i < len(users); i++ {
		userResponse = append(userResponse, users[i].MakeResponse())
	}

	response.Data = userResponse

	responseAPI(response)
}

func validateUser(user models.User) error {
	if user.FullName == "" {
		return errors.New("Full Name is required")
	}

	if user.Password == "" {
		return errors.New("Password is required")
	}

	if user.Email == "" {
		return errors.New("Email is required")
	}

	if user.OfficeID == 0 {
		return errors.New("Office ID is required")
	}

	return nil
}
