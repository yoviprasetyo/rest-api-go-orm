package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"orm/app/models"
	"strconv"

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

	err := c.Bind(&office)
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.ErrorMessage = err.Error()
		fmt.Println(err.Error())

		responseAPI(response)
	}

	errValidation := validateOffice(office)
	if errValidation != nil {
		response.StatusCode = http.StatusBadRequest
		response.ErrorMessage = errValidation.Error()
		fmt.Println(errValidation.Error())
		responseAPI(response)
	}

	models.DB.Create(&office)

	officeResponse = append(officeResponse, office.MakeResponse())

	response.Data = officeResponse

	responseAPI(response)

}

// DeleteOffice controller.
func (controller *Controller) DeleteOffice(c *gin.Context) {

	var (
		office   models.Office
		response = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	id := c.Param("id")
	err := models.DB.First(&office, id).Error

	if err != nil {
		response.ErrorMessage = err.Error()
		response.StatusCode = http.StatusNotFound
		responseAPI(response)
	}

	errDelete := models.DB.Delete(&office, id).Error
	if errDelete != nil {
		// result = responseAPI(nil, errDelete.Error())
		response.ErrorMessage = err.Error()
		response.StatusCode = http.StatusInternalServerError
		responseAPI(response)
	}

	response.Data = office.MakeResponse()
	responseAPI(response)
}

// UpdateOffice controller.
func (controller *Controller) UpdateOffice(c *gin.Context) {

	var (
		office    models.Office
		newOffice models.Office
		response  = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	id := c.Param("id")
	err := c.Bind(&newOffice)
	err = models.DB.First(&office, id).Error
	if err != nil {
		response.ErrorMessage = "Data not found"
		response.StatusCode = http.StatusNotFound
		responseAPI(response)
	}

	errUpdate := models.DB.Model(&office).Updates(newOffice).Error
	if errUpdate != nil {
		response.ErrorMessage = errUpdate.Error()
		response.StatusCode = http.StatusInternalServerError
		responseAPI(response)
	}

	response.Data = newOffice.MakeResponse()

	responseAPI(response)
}

// GetOffice controllers.
func (controller *Controller) GetOffice(c *gin.Context) {
	var (
		officeResponse []gin.H
		offices        []models.Office
		response       = Response{
			Context:       c,
			StatusCode:    http.StatusOK,
			ResponseURL:   ResponseURL{},
			UsePagination: true,
		}
	)

	relation := c.Query("relation")

	page := c.DefaultQuery("page", "1")
	perPage := c.DefaultQuery("per_page", "5")

	perPageInt, _ := strconv.Atoi(perPage)
	pageInt, _ := strconv.Atoi(page)

	models.DB.Limit(perPageInt).Offset(((pageInt - 1) * perPageInt)).Find(&offices)
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

	models.DB.Model(&models.Office{}).Count(&response.Total)

	response.Page = pageInt
	response.PerPage = perPageInt

	nextPageInt := pageInt + 1
	prevPageInt := pageInt - 1

	response.ResponseURL.SetCurrentPageURL("/offices?page=" + page)

	response.ResponseURL.SetFirstPageURL("/offices?page=1")
	response.ResponseURL.SetLastPageURL("/offices?page=" + fmt.Sprint(getLastPage(response)))
	response.ResponseURL.SetNextPageURL("/offices?page=" + strconv.Itoa(nextPageInt))
	response.ResponseURL.SetPrevPageURL("/offices?page=" + strconv.Itoa(prevPageInt))

	responseAPI(response)
}

// GetOneOffice controllers.
func (controller *Controller) GetOneOffice(c *gin.Context) {
	var (
		office   models.Office
		response = Response{
			Context:    c,
			StatusCode: http.StatusOK,
			Total:      1,
		}
	)

	id := c.Param("id")
	err := models.DB.First(&office, id).Error
	if err != nil {
		response.ErrorMessage = "Data Not Found"
		response.StatusCode = http.StatusNotFound
		responseAPI(response)
	}

	response.Data = office.MakeResponse()
	responseAPI(response)
}

// GetSearchOffice controllers.
func (controller *Controller) GetSearchOffice(c *gin.Context) {
	var (
		officeResponse []gin.H
		offices        []models.Office
		response       = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	search := c.Query("search")

	if search == "" {
		response.StatusCode = http.StatusBadRequest
		response.ErrorMessage = "Search query is required"
		responseAPI(response)
	}

	err := models.DB.Where("name LIKE ?", "%"+search+"%").Or("address LIKE ?", "%"+search+"%").Find(&offices).Error
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.ErrorMessage = err.Error()
		responseAPI(response)
	}

	for _, office := range offices {
		officeResponse = append(officeResponse, office.MakeResponse())
	}

	response.Data = officeResponse
	response.Total = int64(len(officeResponse))

	responseAPI(response)
}

func validateOffice(office models.Office) error {
	if office.Name == "" {
		return errors.New("Name is required")
	}

	if office.Address == "" {
		return errors.New("Address is required")
	}

	return nil
}
