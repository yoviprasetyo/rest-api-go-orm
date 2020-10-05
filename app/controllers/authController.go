package controllers

// import (
// 	"fmt"
// 	"net/http"
// 	"orm/app/models"
// 	"os"
// 	"time"

// 	"golang.org/x/crypto/bcrypt"

// 	"github.com/gin-gonic/gin"

// 	jwt "github.com/dgrijalva/jwt-go"
// )

// // Login controllers.
// func (strDB *StrDB) Login(c *gin.Context) {
// 	var (
// 		user   models.User
// 		userDB models.User
// 		result gin.H
// 	)

// 	err := c.Bind(&user)
// 	if err != nil {
// 		// fmt.Println("tidak ada data")
// 	}

// 	errUser := strDB.DB.Where("email = ?", user.Email).First(&userDB).Error
// 	if errUser != nil {
// 		result = gin.H{
// 			"ok":      false,
// 			"message": "User not found",
// 		}
// 	}

// 	var errPass error
// 	if errUser == nil {
// 		errPass := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
// 		if errPass != nil {
// 			result = gin.H{
// 				"ok":      false,
// 				"message": "Password salah",
// 			}
// 		}
// 	}

// 	type authCustomClaims struct {
// 		Email string `json:"email"`
// 		jwt.StandardClaims
// 	}

// 	claims := &authCustomClaims{
// 		userDB.Email,
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
// 			IssuedAt:  time.Now().Unix(),
// 		},
// 	}

// 	fmt.Println(claims)

// 	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)

// 	token, errJwt := sign.SignedString([]byte(os.Getenv("SECRET_TOKEN")))

// 	if errUser == nil && errPass == nil && errJwt == nil {

// 		result = gin.H{
// 			"ok":           true,
// 			"message":      "Success login",
// 			"access_token": token,
// 		}
// 	}

// 	c.JSON(http.StatusOK, result)
// }
