package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"orm/app/helper"
	"orm/app/models"
	"strconv"

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

// DeleteUser controller.
func (controller *Controller) DeleteUser(c *gin.Context) {
	var (
		user     models.User
		response = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	id := c.Param("id")
	err := models.DB.First(&user, id).Error

	if err != nil {
		response.ErrorMessage = err.Error()
		response.StatusCode = http.StatusNotFound
		responseAPI(response)
	}

	errDelete := models.DB.Delete(&user, id).Error
	if errDelete != nil {
		response.ErrorMessage = err.Error()
		response.StatusCode = http.StatusInternalServerError
		responseAPI(response)
	}

	response.Data = user.MakeResponse()
	responseAPI(response)
}

// UpdateUser controller.
func (controller *Controller) UpdateUser(c *gin.Context) {

	var (
		user     models.User
		newUser  models.User
		response = Response{
			Context:    c,
			StatusCode: http.StatusOK,
		}
	)

	id := c.Param("id")
	err := c.Bind(&newUser)
	err = models.DB.First(&user, id).Error
	if err != nil {
		response.ErrorMessage = "Data not found"
		response.StatusCode = http.StatusNotFound
		responseAPI(response)
	}

	errUpdate := models.DB.Model(&user).Updates(newUser).Error
	if errUpdate != nil {
		response.ErrorMessage = errUpdate.Error()
		response.StatusCode = http.StatusInternalServerError
		responseAPI(response)
	}

	response.Data = newUser.MakeResponse()

	responseAPI(response)
}

// GetOneUser controllers.
func (controller *Controller) GetOneUser(c *gin.Context) {
	var (
		user     models.User
		response = Response{
			Context:    c,
			StatusCode: http.StatusOK,
			Total:      1,
		}
	)

	id := c.Param("id")
	err := models.DB.First(&user, id).Error
	if err != nil {
		response.ErrorMessage = "Data Not Found"
		response.StatusCode = http.StatusNotFound
		responseAPI(response)
	}

	response.Data = user.MakeResponse()
	responseAPI(response)
}

// GetUserRedis method.
func (controller *Controller) GetUserRedis(c *gin.Context) {
	var (
		response = Response{}
	)

	relation := c.Query("relation")
	page := c.DefaultQuery("page", "1")
	perPage := c.DefaultQuery("per_page", "5")
	key := "user_" + page + "_" + perPage
	if relation != "" {
		key += "_" + relation
	}

	reply, err := helper.GetRedis(key)

	fmt.Println(key, reply, err)

	if err == nil {
		err = json.Unmarshal(reply, &response)
		if err == nil {
			responseAPI(response)
			return
		}
	}

	controller.GetUser(c)
}

// GetUser controllers.
func (controller *Controller) GetUser(c *gin.Context) {
	var (
		userResponse []gin.H
		users        []models.User
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

	models.DB.Limit(perPageInt).Offset(((pageInt - 1) * perPageInt)).Find(&users)

	for i := 0; i < len(users); i++ {
		switch relation {
		case "office":
			userResponse = append(userResponse, users[i].MakeResponseWithOffice())
		case "todo":
			userResponse = append(userResponse, users[i].MakeResponseWithToDo())
		default:
			userResponse = append(userResponse, users[i].MakeResponse())
		}
	}

	response.Data = userResponse

	key := "user_" + page + "_" + perPage
	if relation != "" {
		key += "_" + relation
	}

	jd, _ := json.Marshal(response)

	reply, err := helper.RedisConn.Do("SET", key, string(jd))
	if err != nil {
		fmt.Println(err.Error())
	}
	if err == nil {
		fmt.Println("Reply", reply)
	}
	reply, err = helper.RedisConn.Do("EXPIRE", key, strconv.Itoa((30 * 60)))
	if err != nil {
		fmt.Println(err.Error())
	}
	if err == nil {
		fmt.Println("Reply", reply)
	}

	models.DB.Model(&models.User{}).Count(&response.Total)

	response.Page = pageInt
	response.PerPage = perPageInt

	nextPageInt := pageInt + 1
	prevPageInt := pageInt - 1

	response.ResponseURL.SetCurrentPageURL("/users?page=" + page)

	response.ResponseURL.SetFirstPageURL("/users?page=1")
	response.ResponseURL.SetLastPageURL("/users?page=" + fmt.Sprint(getLastPage(response)))
	response.ResponseURL.SetNextPageURL("/users?page=" + strconv.Itoa(nextPageInt))
	response.ResponseURL.SetPrevPageURL("/users?page=" + strconv.Itoa(prevPageInt))

	responseAPI(response)

}

// GetSearchUser controllers.
func (controller *Controller) GetSearchUser(c *gin.Context) {
	var (
		userResponse []gin.H
		users        []models.ToDo
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

	err := models.DB.Where("full_name LIKE ?", "%"+search+"%").Or("email LIKE ?", "%"+search+"%").Find(&users).Error
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.ErrorMessage = err.Error()
		responseAPI(response)
	}

	for _, toDo := range users {
		userResponse = append(userResponse, toDo.MakeResponse())
	}

	response.Data = userResponse
	response.Total = int64(len(userResponse))

	responseAPI(response)
}

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
