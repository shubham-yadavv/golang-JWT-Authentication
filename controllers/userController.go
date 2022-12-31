package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shubham-yadavv/golang-JWT-Authentication/config"
	"github.com/shubham-yadavv/golang-JWT-Authentication/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// get email and pass from body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get body",
		})

		return
	}

	// check if user exists
	var existingUser models.User

	config.DB.Where("email = ?", body.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user already exists please login",
		})
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})

		return
	}

	// create the user
	user := models.User{Email: body.Email, Password: string(hashedPassword)}
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})

		return
	}

	// send reponse

	c.JSON(http.StatusOK, gin.H{
		"success": "user created",
		"user":    user,
	})

}

func Login(c *gin.Context) {
	// get email and pass from body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get body",
		})

		return
	}

	// check if user exists
	var user models.User

	config.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user or password",
		})
	}

	// compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user or password",
		})
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
	}

	// send response
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("authtoken", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"success": "user logged in",
	})
}

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("authtoken", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"success": "user logged out",
	})
}

func GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "user profile",
		"user":    c.MustGet("user"),
	})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "token is valid",
	})
}
