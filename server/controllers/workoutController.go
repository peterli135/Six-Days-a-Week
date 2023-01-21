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
		_, httpStatusCode, errorMsg := helper.CheckUserLoggedInCookie(c)
		if errorMsg != "" {
			c.JSON(httpStatusCode, gin.H{"error": errorMsg})
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
		_, httpStatusCode, errorMsg := helper.CheckUserLoggedInCookie(c)
		if errorMsg != "" {
			c.JSON(httpStatusCode, gin.H{"error": errorMsg})
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
		user, httpStatusCode, errorMsg := helper.CheckUserLoggedInCookie(c)
		if errorMsg != "" {
			c.JSON(httpStatusCode, gin.H{"error": errorMsg})
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
		user, httpStatusCode, errorMsg := helper.CheckUserLoggedInCookie(c)
		if errorMsg != "" {
			c.JSON(httpStatusCode, gin.H{"error": errorMsg})
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
		user, httpStatusCode, errorMsg := helper.CheckUserLoggedInCookie(c)
		if errorMsg != "" {
			c.JSON(httpStatusCode, gin.H{"error": errorMsg})
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

		_, httpStatusCode, errorMsg := helper.CheckUserLoggedInCookie(c)
		if errorMsg != "" {
			c.JSON(httpStatusCode, gin.H{"error": errorMsg})
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

func UpdateExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		exerciseID := c.Param("id")

		_, httpStatusCode, errorMsg := helper.CheckUserLoggedInCookie(c)
		if errorMsg != "" {
			c.JSON(httpStatusCode, gin.H{"error": errorMsg})
			return
		}

		var updateExerciseData models.UpdateExercise
		if err := c.BindJSON(&updateExerciseData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(updateExerciseData)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		updatedExercise, err := database.UpdateExercise(exerciseID, updateExerciseData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update exercise."})
			return
		}

		c.JSON(http.StatusOK, updatedExercise)
	}
}

func UpdateWorkout() gin.HandlerFunc {
	return func(c *gin.Context) {
		workoutID := c.Param("id")

		_, httpStatusCode, errorMsg := helper.CheckUserLoggedInCookie(c)
		if errorMsg != "" {
			c.JSON(httpStatusCode, gin.H{"error": errorMsg})
			return
		}

		var updateWorkoutData models.UpdateWorkoutDate
		if err := c.BindJSON(&updateWorkoutData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(updateWorkoutData)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		updatedWorkout, err := database.UpdateWorkout(workoutID, updateWorkoutData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update exercise."})
			return
		}

		c.JSON(http.StatusOK, updatedWorkout)
	}
}
