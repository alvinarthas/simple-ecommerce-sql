package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// IsAuth to identify login or not
// HandlerFunc is to proceed to next function
func IsAuth() gin.HandlerFunc {
	return checkJWT(false, false)
}

// IsAdmin to identify admin or not
func IsAdmin() gin.HandlerFunc {
	return checkJWT(true, false)
}

// HaveStore to identify, the user have a store or not
func HaveStore() gin.HandlerFunc {
	return checkJWT(false, true)
}

func checkJWT(middlewareAdmin bool, checkStore bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		// Check if there is an authorization token
		if len(bearerToken) == 2 {
			// to the callback, providing flexibility.
			token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {

				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

				userRole := bool(claims["user_role"].(bool))
				userStore := bool(claims["user_store"].(bool))
				c.Set("jwt_user_id", claims["user_id"])

				if middlewareAdmin == true && userRole == false {
					c.JSON(403, gin.H{
						"status":  "Forbidden",
						"message": "Only Admin Allowed"})
					c.Abort()
					return
				} else if checkStore == true && userStore == false {
					c.JSON(403, gin.H{
						"status":  "Forbidden",
						"message": "Only Store User Allowed"})
					c.Abort()
					return
				}
			} else {
				c.JSON(422, gin.H{
					"msg": "Invalid Token",
					"err": err,
				})
				c.Abort()
				return
			}
		} else {
			c.JSON(422, gin.H{
				"msg": "Authorization token not provided",
			})
			c.Abort()
			return
		}

	}

}
