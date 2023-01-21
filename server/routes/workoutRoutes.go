package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

// WorkoutRoutes function
func WorkoutRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/api/exercise", controllers.AddExercise())
	incomingRoutes.POST("/api/exercises", controllers.AddMultipleExercises())
	incomingRoutes.POST("/api/exercise/update/:id", controllers.UpdateExercise())
	incomingRoutes.POST("/api/workoutdate", controllers.AddWorkoutDate())
	incomingRoutes.GET("/api/user/workouts", controllers.GetUserWorkouts())
	incomingRoutes.GET("/api/user/workouts/currentyear", controllers.GetUserWorkoutsCurrentYear())
	incomingRoutes.GET("/api/user/workout/:id", controllers.GetExercisesInWorkout())
}
