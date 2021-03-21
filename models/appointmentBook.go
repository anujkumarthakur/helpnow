package models

import (
	"database/sql"
	"fmt"
	"helpnow/config"
	"log"
)

type AppointmentBook struct {
	DoctorId    string `json:"doctor_id"`
	PatientName string `json:"patient_name"`
	Mobile      string `json:"mobile"`
}

func PatientAppointment(appoint AppointmentBook) (*AppointmentBook, error) {
	db := config.GetDB()
	var sehedule SeheduleDoctorTime
	if appoint.DoctorId != "" && appoint.PatientName != "" && appoint.Mobile != "" {
		existId := `SELECT from_time, to_time from doctor_ast where doctor_id=$1;`
		row := db.QueryRow(existId, appoint.DoctorId)
		switch err := row.Scan(&sehedule.AvaliableFromTime, &sehedule.AvaliableToTime); err {
		case sql.ErrNoRows:
			fmt.Println("Doctor Not Exists")
		case nil:
			//strat next process
			fmt.Println(&sehedule.AvaliableFromTime, &sehedule.AvaliableToTime)
		default:
			log.Println(err)
		}
	}
	return
}
