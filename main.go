package main

import (
	"encoding/json"
	"fmt"
	"orm/app/controllers"
	"orm/app/helper"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func main() {
	// models.Migrations(db)

	controller := controllers.Controller{}

	controller.UseDB()

	router := gin.Default()

	initRedis()

	v1 := router.Group("/api/v1")
	{
		{

			v1.GET("/search/office", controller.GetSearchOffice)
			v1.GET("/offices/:id/users", controller.GetUserOffice)
			v1.GET("/offices/:id/show", controller.GetOneOffice)
			v1.POST("/offices", controller.CreateOffice)
			v1.GET("/offices", controller.GetOffice)
			v1.DELETE("/offices/:id", controller.DeleteOffice)
			v1.PUT("/offices/:id", controller.UpdateOffice)
		}
		{
			v1.GET("/search/users", controller.GetSearchUser)
			v1.GET("/users", controller.GetUser)
			v1.GET("/users/:id/show", controller.GetOneUser)
			v1.POST("/users", controller.CreateUser)
			v1.GET("/users/:id/todos", controller.GetToDoUser)
			v1.DELETE("/users/:id", controller.DeleteUser)
			v1.PUT("/users/:id", controller.UpdateUser)
		}
		{
			v1.GET("/search/todos", controller.GetSearchToDo)
			v1.GET("/todos", controller.GetToDo)
			v1.GET("/todos/:id/show", controller.GetOneToDo)
			v1.POST("/todos", controller.CreateToDo)
			v1.DELETE("/todos/:id", controller.DeleteToDo)
			v1.PUT("/todos/:id", controller.UpdateToDo)
		}
	}

	redis := router.Group("/api/v1/redis")
	{
		{
			// v2.GET("/offices/:id/users", controller.GetUserOfficeRedis)
			// v2.GET("/offices/:id/show", controller.GetOneOfficeRedis)
			redis.GET("/offices", controller.GetOfficeRedis)
		}
		{
			redis.GET("/users", controller.GetUserRedis)
			// v2.GET("/users/:id/show", controller.GetOneUserRedis)
			// v2.GET("/users/:id/todos", controller.GetToDoUserRedis)
		}
		{
			redis.GET("/todos", controller.GetToDoRedis)
			// v2.GET("/search/todos", controller.GetSearchToDoRedis)
			// v2.GET("/todos/:id/show", controller.GetOneToDoRedis)
		}
	}
	router.Run()
	// SendRedis()
}

func initRedis() {
	pool := redis.NewPool(
		func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
		10,
	)

	fmt.Println("Connected to redis")

	pool.MaxActive = 10

	helper.RedisConn = pool.Get()
}

// SendRedis method.
func SendRedis() {
	// pool := redis.NewPool(
	// 	func() (redis.Conn, error) {
	// 		return redis.Dial("tcp", "127.0.0.1:6379")
	// 	},
	// 	10,
	// )

	// pool.MaxActive = 10

	// helper.RedisConn = pool.Get()
	// defer helper.RedisConn.Close()

	_, err := helper.RedisConn.Do("HSET", "mahasiswa:1", "name", "Redha Juanda", "nim", "123123123", "ipk", 4.0, "semester", 4)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = helper.RedisConn.Do("SET", "test", "mantapjiwa")
	if err != nil {
		fmt.Println(err.Error())
	}

	reply, err := helper.GetRedis("test")
	if err != nil {
		fmt.Println(err.Error())
	}

	user := map[string]string{
		"name": "Paimo",
		"city": "Surakarta",
	}

	jd, _ := json.Marshal(user)
	fmt.Println(string(jd))

	_, _ = helper.RedisConn.Do("SET", "mahasiswa:1", string(jd))

	reply, err = helper.GetRedis("mahasiswa:1")
	if err == nil {
		fmt.Println(string(reply))
	}

	var newUser map[string]interface{}

	err = json.Unmarshal(reply, &newUser)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(newUser)

	for key, value := range newUser {
		fmt.Println(key, value)
	}

	jd, _ = json.Marshal(newUser)
	fmt.Println(string(jd))

	_, _ = helper.RedisConn.Do("EXPIRE", "mahasiswa:1", "5")
	// helper.RedisConn.Do("DEL", "mahasiswa:1")
}
