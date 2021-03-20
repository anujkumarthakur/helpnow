package controllers

import (
	"helpnow/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Sehedule(c *gin.Context) {
	var sehedule models.DoctorSeheduleTime
	c.BindJSON(&sehedule)
	data, err := models.SeheduleDoctorTime(sehedule)
	c.JSON(http.StatusOK, gin.H{"Error": err, "Data": data})
}
