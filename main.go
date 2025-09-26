package main

import (
	"fmt"
	"net/http"

	"github.com/Pkra99/go-jwt-auth/controllers"
	"github.com/Pkra99/go-jwt-auth/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
}
func main() {
	fmt.Println("Hello world")

	 r := gin.Default()
  
  // Define a simple GET endpoint
  	r.GET("/ping", func(c *gin.Context) {
    	// Return JSON response
    	c.JSON(http.StatusOK, gin.H{
      		"message": "pong",
    	})
  	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)


  
  // Start server on port 8080 (default)
  // Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
  r.Run()
}
