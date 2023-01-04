package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/api/signup", controllers.SignUp())
	incomingRoutes.POST("/api/login", controllers.Login())
	incomingRoutes.GET("/api/user", controllers.User())
	incomingRoutes.POST("/api/logout", controllers.Logout())
	incomingRoutes.POST("/api/exercise", controllers.AddExercise())
	incomingRoutes.POST("/api/exercises", controllers.AddMultipleExercises())
	incomingRoutes.POST("/api/workoutdate", controllers.AddWorkoutDate())
	incomingRoutes.GET("/api/user/workouts", controllers.GetUserWorkouts())
	incomingRoutes.GET("/api/user/workouts/currentyear", controllers.GetUserWorkoutsCurrentYear())
	incomingRoutes.GET("/api/user/workout/:id", controllers.GetExercisesInWorkout())
}
