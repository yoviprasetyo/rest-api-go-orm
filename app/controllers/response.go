package controllers

import "github.com/gin-gonic/gin"

func responseAPI(collection interface{}, count int) gin.H {
	return gin.H{
		"ok":    true,
		"data":  collection,
		"count": count,
	}
}
