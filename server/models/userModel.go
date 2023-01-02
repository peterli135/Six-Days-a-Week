package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct
type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"firstname" validate:"required,min=2,max=100"`
	LastName     *string            `json:"lastname" validate:"required,min=2,max=100"`
	Password     *string            `json:"password" validate:"required,min=6"`
	Email        *string            `json:"email" validate:"email,required"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refreshtoken"`
	CreatedAt    time.Time          `json:"createdat"`
	UpdatedAt    time.Time          `json:"updatedat"`
	UserID       string             `json:"userid"`
}

// Create a subset pattern? Store this months workouts in the user collection, and then another collection will store all the workouts.

// stores one exercise and how many sets/reps you did for that exercise (create a workout if it doesn't exist in database, otherwise, get workout from database)
type Exercise struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   *string            `json:"name" validate:"required"`
	Weight *int               `json:"weight" validate:"required"`
	Sets   *int               `json:"sets" validate:"required"` // figure out a way to change reps (maybe later sets, u did less reps)
	Reps   *int               `json:"reps" validate:"required"`
}

// stores the day of your workout which includes all your exercises you did.
type WorkoutDate struct {
	ID        primitive.ObjectID   `bson:"_id"`
	Date      *string              `json:"date" validate:"required"`
	Name      *string              `json:"name" validate:"required"`
	Exercises []primitive.ObjectID `bson:"exercises"`
	UserID    primitive.ObjectID   `bson:"UserID"`
}

// stores the json data after inserting exercises (because exercises becomes strings, not primitive.ObjectIDs)
type WorkoutDateJSON struct {
	Date      *string  `json:"date" validate:"required"`
	Name      *string  `json:"name" validate:"required"`
	Exercises []string `json:"exercises" validate:"required"`
}

// SignedDetails struct
type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	jwt.RegisteredClaims
}
