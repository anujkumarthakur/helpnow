package controllers

import (
	"helpnow/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Sehedule(c *gin.Context) {
	var sehedule models.DoctorSehedule
	c.BindJSON(&sehedule)
	data, err := models.SeheduleDoctorTime(sehedule)
	c.JSON(http.StatusOK, gin.H{"Error": err, "Data": data})
}

func DoctorsList(c *gin.Context) {
	getDoctorsList, _ := models.GetDoctorsList()
	c.JSON(http.StatusOK, getDoctorsList)
}

func AppointmentBook(c *gin.Context) {
	var appoint models.AppointmentBook
	c.BindJSON(&appoint)
	data, err := models.PatientAppointment(appoint)
	c.JSON(http.StatusOK, gin.H{"Error": err, "Data": data})
}

func CancelApponiment(c *gin.Context) {
	var cancel models.AppointmentCancel
	c.BindJSON(&cancel)
	data, err := models.BookingCancel(cancel)
	c.JSON(http.StatusOK, gin.H{"Error": err, "Data": data})
}
