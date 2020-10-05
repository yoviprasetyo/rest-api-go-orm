package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"orm/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateToDo controller.
func (controller *Controller) CreateToDo(c *gin.Context) {
	var (
		toDoResponse []gin.H
		toDo         models.ToDo
		response     = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	err := c.Bind(&toDo)
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.ErrorMessage = err.Error()
		responseAPI(response)
		return
	}

	errValidation := validateToDo(toDo)
	if errValidation != nil {
		response.StatusCode = http.StatusBadRequest
		response.ErrorMessage = errValidation.Error()
		responseAPI(response)
		return
	}

	models.DB.Create(&toDo)

	response.Data = append(toDoResponse, toDo.MakeResponse())
	responseAPI(response)
}

// DeleteToDo controller.
func (controller *Controller) DeleteToDo(c *gin.Context) {
	var (
		toDo     models.ToDo
		response = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	id := c.Param("id")
	err := models.DB.First(&toDo, id).Error

	if err != nil {
		response.ErrorMessage = err.Error()
		response.StatusCode = http.StatusNotFound
		responseAPI(response)
	}

	errDelete := models.DB.Delete(&toDo, id).Error
	if errDelete != nil {
		response.ErrorMessage = err.Error()
		response.StatusCode = http.StatusInternalServerError
		responseAPI(response)
	}

	response.Data = toDo.MakeResponse()
	responseAPI(response)
}

// UpdateToDo controller.
func (controller *Controller) UpdateToDo(c *gin.Context) {

	var (
		toDo     models.ToDo
		newToDo  models.ToDo
		response = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	id := c.Param("id")
	err := c.Bind(&newToDo)
	err = models.DB.First(&toDo, id).Error
	if err != nil {
		response.ErrorMessage = "Data not found"
		response.StatusCode = http.StatusNotFound
		responseAPI(response)
	}

	errUpdate := models.DB.Model(&toDo).Updates(newToDo).Error
	if errUpdate != nil {
		response.ErrorMessage = errUpdate.Error()
		response.StatusCode = http.StatusInternalServerError
		responseAPI(response)
	}

	response.Data = newToDo.MakeResponse()

	responseAPI(response)
}

// GetToDo controllers.
func (controller *Controller) GetToDo(c *gin.Context) {
	var (
		toDoResponse []gin.H
		toDos        []models.ToDo
		response     = Response{
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

	models.DB.Limit(perPageInt).Offset(((pageInt - 1) * perPageInt)).Find(&toDos)

	for i := 0; i < len(toDos); i++ {
		switch relation {
		case "user":
			toDoResponse = append(toDoResponse, toDos[i].MakeResponseWithUser())
		case "office":
			toDoResponse = append(toDoResponse, toDos[i].MakeResponseWithOffice())
		default:
			toDoResponse = append(toDoResponse, toDos[i].MakeResponse())
		}
	}

	response.Data = toDoResponse

	models.DB.Model(&models.ToDo{}).Count(&response.Total)

	response.Page = pageInt
	response.PerPage = perPageInt

	nextPageInt := pageInt + 1
	prevPageInt := pageInt - 1

	response.ResponseURL.SetCurrentPageURL("/todos?page=" + page)

	response.ResponseURL.SetFirstPageURL("/todos?page=1")
	response.ResponseURL.SetLastPageURL("/todos?page=" + fmt.Sprint(getLastPage(response)))
	response.ResponseURL.SetNextPageURL("/todos?page=" + strconv.Itoa(nextPageInt))
	response.ResponseURL.SetPrevPageURL("/todos?page=" + strconv.Itoa(prevPageInt))

	responseAPI(response)
}

// GetToDoUser controllers.
func (controller *Controller) GetToDoUser(c *gin.Context) {
	var (
		toDoResponse []gin.H
		toDos        []models.ToDo
		response     = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	userID := c.Param("id")
	relation := c.Query("relation")

	models.DB.Where("user_id = ?", userID).Find(&toDos)
	for i := 0; i < len(toDos); i++ {

		switch relation {
		case "user":
			toDoResponse = append(toDoResponse, toDos[i].MakeResponseWithUser())
		case "office":
			toDoResponse = append(toDoResponse, toDos[i].MakeResponseWithOffice())
		default:
			toDoResponse = append(toDoResponse, toDos[i].MakeResponse())
		}
	}

	response.Data = toDoResponse

	responseAPI(response)
}

// GetOneToDo controllers.
func (controller *Controller) GetOneToDo(c *gin.Context) {
	var (
		toDo     models.ToDo
		response = Response{
			Context:    c,
			StatusCode: http.StatusOK,
			Total:      1,
		}
	)

	id := c.Param("id")
	err := models.DB.First(&toDo, id).Error
	if err != nil {
		response.ErrorMessage = "Data Not Found"
		response.StatusCode = http.StatusNotFound
		responseAPI(response)
	}

	response.Data = toDo.MakeResponse()
	responseAPI(response)

}

// GetSearchToDo controllers.
func (controller *Controller) GetSearchToDo(c *gin.Context) {
	var (
		toDoResponse []gin.H
		toDos        []models.ToDo
		response     = Response{
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

	err := models.DB.Where("name LIKE ?", "%"+search+"%").Or("description LIKE ?", "%"+search+"%").Find(&toDos).Error
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.ErrorMessage = err.Error()
		responseAPI(response)
	}

	for _, toDo := range toDos {
		toDoResponse = append(toDoResponse, toDo.MakeResponse())
	}

	response.Data = toDoResponse
	response.Total = int64(len(toDoResponse))

	responseAPI(response)

}

func validateToDo(toDo models.ToDo) error {
	if toDo.Name == "" {
		return errors.New("Name is required")
	}

	if toDo.Description == "" {
		return errors.New("Description is required")
	}

	if toDo.UserID == 0 {
		return errors.New("User ID is required")
	}

	return nil
}
