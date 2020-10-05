package controllers

import (
	"errors"
	"net/http"
	"orm/app/models"

	"github.com/gin-gonic/gin"
)

// CreateOffice controller.
func (controller *Controller) CreateOffice(c *gin.Context) {
	var (
		officeResponse []gin.H
		office         models.Office
		response       = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	controller.UseDB()

	err := c.Bind(&office)
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.ErrorMessage = err.Error()

		responseAPI(response)
		return
	}

	errValidation := validateOffice(office)
	if errValidation != nil {
		response.StatusCode = http.StatusBadRequest
		response.ErrorMessage = errValidation.Error()
		responseAPI(response)
		return
	}

	models.DB.Create(&office)

	officeResponse = append(officeResponse, office.MakeResponse())

	response.Data = officeResponse
	responseAPI(response)
}

// // DeleteOffice controller.
// func (controller *Controller) DeleteOffice(c *gin.Context) {
// 	var (
// 		office     models.Office
// 		result     gin.H
// 		resultCode int
// 	)

// 	id := c.Param("id")
// 	err := controller.DB.DB.First(&office, id).Error
// 	resultCode = http.StatusOK
// 	result = gin.H{
// 		"ok":      true,
// 		"message": "Data deleted",
// 	}
// 	if err != nil {
// 		result = gin.H{
// 			"ok":      false,
// 			"message": "Data not found",
// 		}
// 		resultCode = http.StatusNotFound
// 	}

// 	errDelete := controller.DB.DB.Delete(&office, id).Error
// 	if errDelete != nil {
// 		result = responseAPI(nil, errDelete.Error())
// 		resultCode = http.StatusInternalServerError
// 	}

// 	c.JSON(resultCode, result)
// }

// // UpdateOffice controller.
// func (controller *Controller) UpdateOffice(c *gin.Context) {

// 	var (
// 		office     models.Office
// 		newOffice  models.Office
// 		result     gin.H
// 		resultCode = http.StatusOK
// 	)

// 	result = gin.H{
// 		"ok":      true,
// 		"message": "Success update",
// 	}

// 	result = responseAPI()

// 	id := c.Param("id")
// 	err := c.Bind(&newOffice)
// 	err = controller.DB.DB.First(&office, id).Error
// 	if err != nil {
// 		result = responseAPI(nil, "Data not found")
// 		resultCode = http.StatusNotFound
// 	}

// 	if resultCode != http.StatusNotFound {
// 		errUpdate := controller.DB.DB.Model(&office).Updates(newOffice).Error
// 		if errUpdate != nil {
// 			result = responseAPI(nil, errUpdate.Error())
// 			resultCode = http.StatusInternalServerError
// 		}
// 	}

// 	c.JSON(resultCode, result)
// }

// GetOffice controllers.
func (controller *Controller) GetOffice(c *gin.Context) {
	var (
		officeResponse []gin.H
		offices        []models.Office
		response       = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	controller.UseDB()

	relation := c.Query("relation")

	models.DB.Find(&offices)
	for i := 0; i < len(offices); i++ {
		switch relation {
		case "user":
			officeResponse = append(officeResponse, offices[i].MakeResponseWithUser())
		case "todo":
			officeResponse = append(officeResponse, offices[i].MakeResponseWithToDo())
		default:
			officeResponse = append(officeResponse, offices[i].MakeResponse())
		}

	}

	response.Data = officeResponse

	responseAPI(response)
}

// // GetOneOffice controllers.
// func (controller *Controller) GetOneOffice(c *gin.Context) {
// 	var (
// 		office     models.Office
// 		resultCode = http.StatusOK
// 		result     gin.H
// 	)

// 	id := c.Param("id")
// 	err := controller.DB.DB.First(&office, id).Error
// 	if err != nil {
// 		resultCode = http.StatusNotFound
// 		result = gin.H{
// 			"ok":      false,
// 			"message": "Data not found",
// 		}
// 	}

// 	if resultCode == http.StatusOK {
// 		result = responseAPI(office, "")
// 	}

// 	c.JSON(resultCode, result)

// }

// // GetSearchOffice controllers.
// func (controller *Controller) GetSearchOffice(c *gin.Context) {
// 	var (
// 		office     []models.Office
// 		resultCode = http.StatusOK
// 		result     gin.H
// 	)

// 	search := c.Query("search")
// 	strDBQuery := controller.DB.DB

// 	if search == "" {
// 		resultCode = http.StatusBadRequest
// 		result = gin.H{
// 			"ok":      false,
// 			"message": "Search keyword is required",
// 		}
// 	}

// 	if search != "" {
// 		err := strDBQuery.Where("name LIKE ?", "%"+search+"%").Or("address LIKE ?", "%"+search+"%").Find(&office).Error
// 		if err != nil {
// 			resultCode = http.StatusNotFound
// 			result = gin.H{
// 				"ok":      false,
// 				"message": err.Error,
// 			}
// 		}
// 	}

// 	if resultCode == http.StatusOK {
// 		result = responseAPI(office, "")
// 	}

// 	c.JSON(resultCode, result)

// }

func validateOffice(office models.Office) error {
	if office.Name == "" {
		return errors.New("Name is required")
	}

	if office.Address == "" {
		return errors.New("Address is required")
	}

	return nil
}
