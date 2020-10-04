package helper

import "golang.org/x/crypto/bcrypt"

// YMDHIS time format.
var YMDHIS = "2006-01-02 15:04:05"

// HashPassword function.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
