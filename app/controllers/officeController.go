package controllers

import (
	"fmt"
	"net/http"
	"orm/app/models"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateOffice controller.
func (strDB *StrDB) CreateOffice(c *gin.Context) {
	var (
		office models.Office
		result gin.H
	)

	err := c.Bind(&office)
	if err != nil {
		fmt.Println("tidak ada data")
	}

	strDB.DB.Create(&office)
	result = gin.H{
		"result": gin.H{
			"name": office.Name,
			"time": time.Now(),
		},
	}

	c.JSON(http.StatusOK, result)

}

// DeleteOffice controller.
func (strDB *StrDB) DeleteOffice(c *gin.Context) {
	var (
		office     models.Office
		result     gin.H
		resultCode int
	)

	id := c.Param("id")
	err := strDB.DB.First(&office, id).Error
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

	errDelete := strDB.DB.Delete(&office, id).Error
	if errDelete != nil {
		result = gin.H{
			"ok":      false,
			"message": errDelete.Error,
		}
		resultCode = http.StatusInternalServerError
	}

	c.JSON(resultCode, result)
}

// UpdateOffice controller.
func (strDB StrDB) UpdateOffice(c *gin.Context) {

	var (
		office     models.Office
		newOffice  models.Office
		result     gin.H
		resultCode = http.StatusOK
	)

	result = gin.H{
		"ok":      true,
		"message": "Success update",
	}

	id := c.Param("id")
	err := c.Bind(&newOffice)
	err = strDB.DB.First(&office, id).Error
	if err != nil {
		result = gin.H{
			"ok":      false,
			"message": "Data not found",
		}
		resultCode = http.StatusNotFound
	}

	if resultCode != http.StatusNotFound {
		errUpdate := strDB.DB.Model(&office).Updates(newOffice).Error
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

// GetOffice controllers.
func (strDB *StrDB) GetOffice(c *gin.Context) {
	var (
		office []models.Office
		result gin.H
	)

	strDB.DB.Find(&office)

	if len(office) <= 0 {
		arrayNil := []models.Office{}
		result = gin.H{
			"result": arrayNil,
			"count":  0,
		}
	}

	if len(office) > 0 {
		result = responseAPI(office, len(office))
	}

	c.JSON(http.StatusOK, result)
}

// GetOneOffice controllers.
func (strDB *StrDB) GetOneOffice(c *gin.Context) {
	var (
		office     models.Office
		resultCode = http.StatusOK
		result     gin.H
	)

	id := c.Param("id")
	err := strDB.DB.First(&office, id).Error
	if err != nil {
		resultCode = http.StatusNotFound
		result = gin.H{
			"ok":      false,
			"message": "Data not found",
		}
	}

	if resultCode == http.StatusOK {
		result = responseAPI(office, 1)
	}

	c.JSON(resultCode, result)

}

// GetSearchOffice controllers.
func (strDB *StrDB) GetSearchOffice(c *gin.Context) {
	var (
		office     []models.Office
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
		err := strDBQuery.Where("name LIKE ?", "%"+search+"%").Or("address LIKE ?", "%"+search+"%").Find(&office).Error
		if err != nil {
			resultCode = http.StatusNotFound
			result = gin.H{
				"ok":      false,
				"message": err.Error,
			}
		}
	}

	if resultCode == http.StatusOK {
		result = responseAPI(office, len(office))
	}

	c.JSON(resultCode, result)

}
