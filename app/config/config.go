package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect the DB.
func Connect() *gorm.DB {
	err := godotenv.Load(".env")
	var userDatabase, passDatabase, portDatabase, hostDatabase, nameDatabase string

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	userDatabase = os.Getenv("USER_DATABASE")
	passDatabase = os.Getenv("PASS_DATABASE")
	portDatabase = os.Getenv("PORT_DATABASE")
	hostDatabase = os.Getenv("HOST_DATABASE")
	nameDatabase = os.Getenv("NAME_DATABASE")

	conn := userDatabase + ":" + passDatabase + "@tcp(" + hostDatabase + ":" + portDatabase + ")/" + nameDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, errConn := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if errConn != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Koneksi sukses")
	return db
}
