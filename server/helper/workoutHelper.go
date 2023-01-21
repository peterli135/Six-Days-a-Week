package helper

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

// function to check if there is a user logged in (cookie must be valid)
func CheckUserLoggedInCookie(c *gin.Context) (user models.User, httpStatusCode int, message string) {
	// check if there is a valid cookie (meaning user is logged in)
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return models.User{}, http.StatusInternalServerError, "No valid cookie."
	}
	user, msg := AuthenticateUser(cookie)
	if msg != "" {
		return user, http.StatusUnauthorized, msg
	}
	return user, 0, ""
}
