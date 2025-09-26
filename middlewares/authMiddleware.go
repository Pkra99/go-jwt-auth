package middlewares

import "github.com/gin-gonic/gin"


func authMiddleware(c *gin.Context){

	// get the cookie from req
	// tokenString, err := c.Cookie("Authorization")


	// decode and validate the token

	// check for token expiry

	// find the user with the token

	// attach to req

	// continue
	c.Next()
}