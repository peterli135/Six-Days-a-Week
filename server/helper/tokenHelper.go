package helper

import (
	"log"
	"os"
	"time"

	"server/database"
	"server/models"

	"github.com/golang-jwt/jwt/v4"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

// GenerateAllTokens function generates both the detailed token and refresh token
func GenerateAllTokens(email string, firstName string, lastName string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := models.SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Uid:       uid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    uid,
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(24))), // expires in 1 day
		},
	}

	refreshClaims := models.SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    uid,
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(24))),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

// ValidateToken function validates the JWT token
func ValidateToken(signedToken string) (claims *models.SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		models.SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*models.SignedDetails)
	if !ok {
		msg = "The token is invalid."
		msg = err.Error()
		return
	}

	if claims.ExpiresAt.Time.After(time.Now().Local()) {
		msg = "Token is expired."
		msg = err.Error()
		return
	}

	return claims, msg
}

// function that authenticates the user based on the cookie given
func AuthenticateUser(storedCookie string) (userAccount models.User, msg string) {
	token, err := jwt.ParseWithClaims(storedCookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		msg = "Unauthenticated."
		return
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	user, err := database.GetUserByUserID(claims.Issuer)
	if err != nil {
		msg = "Unable to find account."
	}

	return user, msg
}
