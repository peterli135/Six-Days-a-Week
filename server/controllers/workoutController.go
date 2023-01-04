package controllers

import (
	"net/http"
	"server/database"
	"server/helper"
	"server/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check if there is a valid cookie (meaning user is logged in)
		_, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No valid cookie."})
			return
		}

		var exercise models.Exercise
		if err := c.BindJSON(&exercise); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(exercise)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// find if an existing exercise exists. if it doesn't, then create new exercise and add to database
		exerciseFound, err := database.GetExerciseByName(exercise.Name, exercise.Weight, exercise.Sets, exercise.Reps)
		if err != nil {
			exercise.ID = primitive.NewObjectID()

			resultExerciseInsert, insertErr := database.InsertExercise(&exercise)
			if insertErr != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "This workout exercise was unable to be created."})
				return
			}

			c.JSON(http.StatusOK, resultExerciseInsert)
			return
		}

		c.JSON(http.StatusOK, exerciseFound.ID)
	}
}

func AddMultipleExercises() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check if there is a valid cookie (meaning user is logged in)
		_, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No valid cookie."})
			return
		}

		var exercises []models.Exercise
		if err := c.BindJSON(&exercises); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Var(exercises, "dive")
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		insertedExerciseIDs, err := database.InsertExercises(exercises)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Insert exercises failed."})
			return
		}

		c.JSON(http.StatusOK, insertedExerciseIDs)
	}
}

func AddWorkoutDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check if there is a valid cookie (meaning user is logged in)
		cookie, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No valid cookie."})
			return
		}
		user, msg := helper.AuthenticateUser(cookie)
		if msg != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}

		var workoutDateJSON models.WorkoutDateJSON
		if err := c.BindJSON(&workoutDateJSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(workoutDateJSON)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		var primitiveExerciseIDs []primitive.ObjectID
		for _, exerciseID := range workoutDateJSON.Exercises {
			primitiveExerciseID, pErr := primitive.ObjectIDFromHex(exerciseID)
			if pErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": pErr})
				return
			}
			primitiveExerciseIDs = append(primitiveExerciseIDs, primitiveExerciseID)
		}

		var workoutDate models.WorkoutDate
		workoutDate.ID = primitive.NewObjectID()
		workoutDate.UserID = user.ID
		workoutDate.Name = workoutDateJSON.Name
		workoutDate.Date = workoutDateJSON.Date
		workoutDate.Exercises = primitiveExerciseIDs

		resultWorkoutDateInsert, insertErr := database.InsertWorkoutDate(&workoutDate)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This workout day was unable to be created."})
			return
		}

		c.JSON(http.StatusOK, resultWorkoutDateInsert)
	}
}

func GetUserWorkouts() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check if there is a valid cookie (meaning user is logged in)
		cookie, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No valid cookie."})
			return
		}
		user, msg := helper.AuthenticateUser(cookie)
		if msg != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}

		workouts, err := database.GetWorkoutsByUser(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, workouts)
	}
}

func GetUserWorkoutsCurrentYear() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check if there is a valid cookie (meaning user is logged in)
		cookie, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No valid cookie."})
			return
		}
		user, msg := helper.AuthenticateUser(cookie)
		if msg != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}

		currentYearWorkouts, err := database.GetCurrentYearWorkoutsByUser(user.ID, strconv.Itoa(time.Now().Year()))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, currentYearWorkouts)
	}
}

func GetExercisesInWorkout() gin.HandlerFunc {
	return func(c *gin.Context) {
		workoutID := c.Param("id")

		// check if there is a valid cookie (meaning user is logged in)
		cookie, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No valid cookie."})
			return
		}
		_, msg := helper.AuthenticateUser(cookie)
		if msg != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}

		workout, err := database.GetWorkoutByID(workoutID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No workout found."})
			return
		}

		exercises, err := database.GetExercisesByWorkoutID(workout)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find one of the exercises in the database."})
			return
		}

		c.JSON(http.StatusOK, exercises)
	}
}

//TODO: how to get exercises from the GetUserWorkouts since they are in a ID form...
// maybe go through each workout and then find the workout by the id and then replace the id with the actual workout information, then return workout

// also, later on in the road, if want to have the ability to allow users to add previous workout day items, how to account for that in the addworkout function?
// probably need to account for it in the frontend: for ex, if user chooses to add a previous workout day, have that request be different from regularly adding a workout
