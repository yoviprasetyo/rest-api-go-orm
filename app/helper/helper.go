package helper

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"golang.org/x/crypto/bcrypt"
)

// YMDHIS time format.
var YMDHIS = "2006-01-02 15:04:05"

// AppURL determine the URL of the application.
var AppURL string

// RedisConn determine the redis Pool.
var RedisConn redis.Conn

// HashPassword function.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// BaseURL return the AppURL and given path.
func BaseURL(path string) string {
	return AppURL + path
}

// GetRedis value.
func GetRedis(key string) ([]byte, error) {
	fmt.Println(RedisConn)
	reply, err := redis.Bytes(RedisConn.Do("GET", key))
	return reply, err
}
