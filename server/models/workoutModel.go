package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Database schema design: Embed created workouts with userIDs and exerciseIDs to save storage?

// struct for creating/storing exercise names (ex. bench press, deadlift, squat, etc.)
type ExerciseName struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name *string            `json:"name" validate:"required"`
}

// stores one exercise and how many sets/reps you did for that exercise (create a workout if it doesn't exist in database, otherwise, get workout from database)
type Exercise struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   *string            `json:"name" validate:"required"`
	Weight *int               `json:"weight" validate:"required"`
	Sets   *int               `json:"sets" validate:"required"` // figure out a way to change reps (maybe later sets, u did less reps)
	Reps   *int               `json:"reps" validate:"required"`
}

// store the updated exercise data in this struct
type UpdateExercise struct {
	Name   *string `json:"name" validate:"required"`
	Weight *int    `json:"weight" validate:"required"`
	Sets   *int    `json:"sets" validate:"required"`
	Reps   *int    `json:"reps" validate:"required"`
}

// stores the day of your workout which includes all your exercises you did.
type WorkoutDate struct {
	ID        primitive.ObjectID   `bson:"_id"`
	Date      *string              `json:"date" validate:"required"`
	Name      *string              `json:"name" validate:"required"`
	Exercises []primitive.ObjectID `bson:"exercises"`
	UserID    primitive.ObjectID   `bson:"UserID"`
}

type UpdateWorkoutDate struct {
	Date      *string              `json:"date" validate:"required"`
	Name      *string              `json:"name" validate:"required"`
	Exercises []primitive.ObjectID `bson:"exercises"`
}

// stores the json data after inserting exercises (because exercises becomes strings, not primitive.ObjectIDs)
type WorkoutDateJSON struct {
	Date      *string  `json:"date" validate:"required"`
	Name      *string  `json:"name" validate:"required"`
	Exercises []string `json:"exercises" validate:"required"`
}
