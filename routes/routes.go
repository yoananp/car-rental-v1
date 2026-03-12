package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yoananp/car-rental-v1/controllers"
)

func Route() *gin.Engine {
	router := gin.Default()

	// Customer endpoints
	router.POST("/customers", controllers.InsertCustomer)
	router.GET("/customers", controllers.GetAllCustomer)
	router.PUT("/customers/:id", controllers.UpdateCustomer)
	router.DELETE("/customers/:id", controllers.DeleteCustomer)

	// Cars endpoints
	router.POST("/cars", controllers.InsertCar)
	router.GET("/cars", controllers.GetAllCar)
	router.PUT("/cars/:id", controllers.UpdateCar)
	router.DELETE("/cars/:id", controllers.DeleteCar)

	// Bookings endpoints
	router.POST("/bookings", controllers.InsertBooking)
	router.GET("/bookings", controllers.GetAllBooking)
	router.PUT("/bookings/:id", controllers.UpdateBooking)
	router.DELETE("/bookings/:id", controllers.DeleteBooking)

	return router
}
