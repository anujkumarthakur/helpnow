package routers

import (
	control "helpnow/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/schedule", control.Sehedule)
	router.GET("/list", control.DoctorsList)
	router.POST("/book", control.AppointmentBook)
	return router
}
