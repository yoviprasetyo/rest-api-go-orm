package controllers

import (
	"errors"
	"net/http"
	"orm/app/models"

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

	controller.UseDB()

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

/*
// DeleteToDo controller.
func (controller *Controller) DeleteToDo(c *gin.Context) {
	var (
		todo       models.ToDo
		result     gin.H
		resultCode int
	)

	id := c.Param("id")
	err := controller.DB.DB.First(&todo, id).Error
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

	errDelete := controller.DB.DB.Delete(&todo, id).Error
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
func (controller *Controller) UpdateToDo(c *gin.Context) {

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
	err = controller.DB.DB.First(&todo, id).Error
	if err != nil {
		result = gin.H{
			"ok":      false,
			"message": "Data not found",
		}
		resultCode = http.StatusNotFound
	}

	if resultCode != http.StatusNotFound {
		errUpdate := controller.DB.DB.Model(&todo).Updates(newToDo).Error
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
*/
// GetToDo controllers.
func (controller *Controller) GetToDo(c *gin.Context) {
	var (
		toDoResponse []gin.H
		toDos        []models.ToDo
		response     = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	controller.UseDB()

	models.DB.Find(&toDos)
	for i := 0; i < len(toDos); i++ {
		toDoResponse = append(toDoResponse, toDos[i].MakeResponseWithUser())
	}

	response.Data = toDoResponse

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

	controller.UseDB()

	userID := c.Param("id")
	upper := c.Query("upper")

	models.DB.Where("user_id = ?", userID).Find(&toDos)
	for i := 0; i < len(toDos); i++ {

		switch upper {
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

/*
// GetOneToDo controllers.
func (controller *Controller) GetOneToDo(c *gin.Context) {
	var (
		todo       models.ToDo
		resultCode = http.StatusOK
		result     gin.H
	)

	id := c.Param("id")
	err := controller.DB.DB.First(&todo, id).Error
	if err != nil {
		resultCode = http.StatusNotFound
		result = gin.H{
			"ok":      false,
			"message": "Data not found",
		}
	}

	if resultCode == http.StatusOK {
		result = responseAPI(todo, "")
	}

	c.JSON(resultCode, result)

}

// GetSearchToDo controllers.
func (controller *Controller) GetSearchToDo(c *gin.Context) {
	var (
		todo       []models.ToDo
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
		result = responseAPI(todo, "")
	}

	c.JSON(resultCode, result)

}
*/

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
