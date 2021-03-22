package models

import (
	"fmt"
	"helpnow/config"
)

type AppointmentCancel struct {
	DoctorId  string `json:"doctor_id"`
	PatientId string `json:"patient_id"`
}

func BookingCancel(cancel AppointmentCancel) (string, error) {
	db := config.GetDB()
	var err error
	if cancel.DoctorId != "" && cancel.PatientId != "" {
		cancelQuery := `SELECT doctor_id from book_ast where appoiment_id= $1;`
		row := db.QueryRow(cancelQuery, cancel.PatientId)
		err = row.Scan(&cancel.DoctorId, &cancel.PatientId)
		if err != nil {
			fmt.Println(err)
		}
		sqlDelete := `DELETE FROM book_ast WHERE appoiment_id= $1;`
		_, err = db.Exec(sqlDelete, 1)
		if err != nil {
			fmt.Println(err)
		}
	}

	// res := AppointmentCancel{
	// 	DoctorId:  cancel.DoctorId,
	// 	PatientId: cancel.PatientId}
	return "Your Apponiment Cancel Sucessfully", nil
}
