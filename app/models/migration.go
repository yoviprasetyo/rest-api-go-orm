package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations function.
func Migrations(db *gorm.DB) {
	var check bool
	check = db.Migrator().HasTable(&Office{})
	if !check {
		db.Migrator().CreateTable(&Office{})
		fmt.Println("Create table offices")
	}
	check = db.Migrator().HasTable(&User{})
	if !check {
		db.Migrator().CreateTable(&User{})
		fmt.Println("Create table users")
	}
	check = db.Migrator().HasTable(&ToDo{})
	if !check {
		db.Migrator().CreateTable(&ToDo{})
		fmt.Println("Create table to_dos")
	}
}
