package controllers

import (
	"net/http"
	"os"
	"time"
	"github.com/Pkra99/go-jwt-auth/initializers"
	"github.com/Pkra99/go-jwt-auth/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup (c *gin.Context) {
	// get user credentails from req body
	var body struct {
		Email 	 string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// create user and save to db
	user := models.User{Email: body.Email, Password:  string(hashedPassword)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"message" : "User signedup sucessfully",
	})
}

func Login(c *gin.Context) {
	//get user credentials from req body
	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// find user in db
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// generate a jwt token (access token)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	//sign and get encoded token as string
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// send back response
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "`", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

func GenerateAccessToken(userID uint, email string) (string, error){
	claims := jwt.MapClaims{
		"sub": userID,
		"emial": email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("SECRET")
	return token.SignedString([]byte(secret))
}

func Validate (c *gin.Context) {

	u, _ := c.Get("user")
	user := u.(models.User)

	c.JSON(http.StatusOK, gin.H{
		"message": user.Email,
	})
}