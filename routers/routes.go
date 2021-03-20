package routers

import (
	control "helpnow/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/schedule", control.Sehedule)
	return router
}
