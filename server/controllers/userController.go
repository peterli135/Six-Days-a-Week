package controllers

import (
	"context"
	"log"
	"net/http"
	"server/database"
	"server/helper"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

// HashPassword function used to encrypt the password before it gets stored in the database.
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

// VerifyPassword function checks the input password and verifies it's the password in the database.
func VerifyPassword(userPassword string, databasePassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = "Login or password is incorrect."
		check = false
	}
	return check, msg
}

// CreateUser is the API used to get a single user.
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := database.GetUserCountByEmail(*user.Email)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while checking for the email."})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This email already exists."})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserID = user.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, user.UserID)
		user.Token = &token
		user.RefreshToken = &refreshToken

		resultInsertionNumber, insertErr := database.InsertUser(&user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User item was unable to be created."})
			return
		}

		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

// Login is the API used to get a single user
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userFound, err := database.GetUserByEmail(ctx, user.Email)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No user associated with the email provided."})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *userFound.Password)
		defer cancel()
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(*userFound.Email, *userFound.FirstName, *userFound.LastName, userFound.UserID)

		updatedUser, err := database.UpdateAllTokens(token, refreshToken, userFound.UserID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"data": "Error occurred while logging in.",
			})
		}

		storeCookie := &http.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24),
			Path:     "/",
			HttpOnly: true,
		}

		http.SetCookie(c.Writer, storeCookie)

		c.JSON(http.StatusOK, gin.H{"message": "success", "object": updatedUser})
	}
}

// function to read the session cookie and then return the user that is found from the token
func User() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No valid cookie."})
		}

		user, msg := helper.AuthenticateUser(cookie)
		if msg != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		}

		c.JSON(http.StatusOK, user)
	}
}

// function to logout the user
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		storeCookie := &http.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			Path:     "/",
			HttpOnly: true,
		}

		http.SetCookie(c.Writer, storeCookie)

		c.JSON(http.StatusOK, gin.H{"message": "successfully logged out."})
	}
}
