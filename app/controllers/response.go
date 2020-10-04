package controllers

import "github.com/gin-gonic/gin"

func responseAPI(response Response) {
	var (
		result = gin.H{
			"status": "success",
		}
	)
	if response.ErrorMessage != "" {
		result["status"] = "error"
		result["message"] = response.ErrorMessage
	}
	if response.ErrorMessage == "" {
		result["data"] = response.Data
	}

	response.Context.JSON(response.StatusCode, result)
}

// Response struct.
type Response struct {
	Context      *gin.Context
	StatusCode   int
	Data         interface{}
	ErrorMessage string
}
