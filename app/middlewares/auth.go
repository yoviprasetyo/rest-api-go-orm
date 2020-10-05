package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Auth middleware.
func Auth(c *gin.Context) {
	// err := godotenv.Load(".env")
	// var secret string

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	// secret = os.Getenv("SECRET_TOKEN")
	// tokenString := c.GetHeader("Authorization")

	// fmt.Println("Token String", tokenString)

	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	if jwt.GetSigningMethod("HS256") != token.Method {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}

	// 	return []byte(secret), nil
	// })

	// fmt.Println(token, err)

	// if token == nil || err != nil {

	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"ok":      false,
	// 		"message": err.Error(),
	// 	})
	// 	c.Abort()
	// }
}
