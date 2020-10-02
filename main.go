package main

import (
	"orm/app/config"
	"orm/app/controllers"
	"orm/app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.Connect()
	models.Migrations(db)

	StrDB := controllers.StrDB{DB: db}

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		{
			v1.GET("/offices", StrDB.GetOffice)
			v1.GET("/search/office", StrDB.GetSearchOffice)
			v1.GET("/offices/:id/show", StrDB.GetOneOffice)
			v1.POST("/offices", StrDB.CreateOffice)
			v1.DELETE("/offices/:id", StrDB.DeleteOffice)
			v1.PUT("/offices/:id", StrDB.UpdateOffice)
		}
		{
			v1.GET("/search/users", StrDB.GetSearchUser)
			v1.GET("/users", StrDB.GetUser)
			v1.GET("/users/:id/show", StrDB.GetOneUser)
			v1.POST("/users", StrDB.CreateUser)
			v1.DELETE("/users/:id", StrDB.DeleteUser)
			v1.PUT("/users/:id", StrDB.UpdateUser)
		}
		{
			v1.GET("/search/to_dos", StrDB.GetSearchToDo)
			v1.GET("/to_dos", StrDB.GetToDo)
			v1.GET("/to_dos/:id/show", StrDB.GetOneToDo)
			v1.POST("/to_dos", StrDB.CreateToDo)
			v1.DELETE("/to_dos/:id", StrDB.DeleteToDo)
			v1.PUT("/to_dos/:id", StrDB.UpdateToDo)
		}
	}

	router.Run()
}
