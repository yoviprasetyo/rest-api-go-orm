package controllers

import (
	"fmt"
	"net/http"
	"orm/app/models"

	"github.com/gin-gonic/gin"
)

// CreateOffice controller.
func (controller *Controller) CreateOffice(c *gin.Context) {
	var (
		officeResponse []models.OfficeResponse
		office         models.Office
		response       = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	err := c.Bind(&office)
	if err != nil {
		fmt.Println("tidak ada data")
	}

	controller.DB.DB.Create(&office)

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

// // GetOffice controllers.
// func (controller *Controller) GetOffice(c *gin.Context) {
// 	var (
// 		office []models.Office
// 		result gin.H
// 	)

// 	controller.DB.DB.Find(&office)

// 	if len(office) <= 0 {
// 		arrayNil := []models.Office{}
// 		result = gin.H{
// 			"result": arrayNil,
// 			"count":  0,
// 		}
// 	}

// 	if len(office) > 0 {
// 		result = responseAPI(office, "")
// 	}

// 	c.JSON(http.StatusOK, result)
// }

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
