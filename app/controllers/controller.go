package controllers

import (
	"orm/app/config"
	"orm/app/models"
)

// Controller struct.
type Controller struct {
}

// UseDB in the controller.
func (controller *Controller) UseDB() {
	models.DB = config.Connect()
}
