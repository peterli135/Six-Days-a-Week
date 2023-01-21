package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// struct for users
type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"firstname" validate:"required,min=2,max=100"`
	LastName     *string            `json:"lastname" validate:"required,min=2,max=100"`
	Password     *string            `json:"password" validate:"required,min=6"` //`json:"-" validate:"required,min=6"`
	Email        *string            `json:"email" validate:"email,required"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refreshtoken"`
	CreatedAt    time.Time          `json:"createdat"`
	UpdatedAt    time.Time          `json:"updatedat"`
	UserID       string             `json:"userid"`
}

// SignedDetails struct
type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	jwt.RegisteredClaims
}
