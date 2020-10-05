package main

import (
	"orm/app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// models.Migrations(db)

	controller := controllers.Controller{}

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		{

			// v1.GET("/search/office", controller.GetSearchOffice)
			v1.GET("/offices/:id/users", controller.GetUserOffice)
			v1.POST("/offices", controller.CreateOffice)
			v1.GET("/offices", controller.GetOffice)
			// v1.DELETE("/offices/:id", controller.DeleteOffice)
			// v1.PUT("/offices/:id", controller.UpdateOffice)
		}
		{
			// 	v1.GET("/search/users", controller.GetSearchUser)
			// 	v1.GET("/users", controller.GetUser)
			// 	v1.GET("/users/:id/show", controller.GetOneUser)
			v1.POST("/users", controller.CreateUser)
			v1.GET("/users/:id/todos", controller.GetToDoUser)
			// 	v1.DELETE("/users/:id", controller.DeleteUser)
			// 	v1.PUT("/users/:id", controller.UpdateUser)
		}
		{
			// 	v1.GET("/search/to_dos", controller.GetSearchToDo)
			v1.GET("/todos", controller.GetToDo)

			// 	v1.GET("/to_dos/:id/show", controller.GetOneToDo)
			v1.POST("/todos", controller.CreateToDo)
			// 	v1.DELETE("/to_dos/:id", controller.DeleteToDo)
			// 	v1.PUT("/to_dos/:id", controller.UpdateToDo)
		}
	}

	router.Run()
}
