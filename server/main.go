package main

import (
	"os"

	"server/middleware"
	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		//AllowAllOrigins:  true,
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin, X-Requested-With, Content-Type, Accept"},
	}))

	/*
		// these are the endpoints
		//C
		router.POST("/order/create", routes.AddOrder)
		//R
		router.GET("/waiter/:waiter", routes.GetOrdersByWaiter)
		router.GET("/orders", routes.GetOrders)
		router.GET("/order/:id/", routes.GetOrderById)
		//U
		router.PUT("/waiter/update/:id", routes.UpdateWaiter)
		router.PUT("/order/update/:id", routes.UpdateOrder)
		//D
		router.DELETE("/order/delete/:id", routes.DeleteOrder)
	*/

	routes.UserRoutes(router)
	routes.WorkoutRoutes(router)

	router.Use(middleware.Authentication())

	// API Dummy Endpoints (Test)
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for API-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for API-2"})
	})

	//this runs the server and allows it to listen to requests.
	router.Run(":" + port)
}
