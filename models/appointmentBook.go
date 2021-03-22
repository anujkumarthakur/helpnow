package models

import (
	"fmt"
	"helpnow/config"
)

type AppointmentBook struct {
	DoctorId    string `json:"doctor_id"`
	PatientName string `json:"patient_name"`
	Mobile      string `json:"mobile"`
}

type BookSlot struct {
	DoctorId string `json:"doctor_id"`
	BookSlot string `json:"book_slot"`
}

type DoctorApponimntBooked struct {
	DoctorDetail DoctorSehedule `json:"doctor_detail"`
	SlotBooked   BookSlot       `json:"slot_booked"`
}

func PatientAppointment(appoint AppointmentBook) (*DoctorApponimntBooked, error) {
	db := config.GetDB()
	var sehedule DoctorSehedule
	var book BookSlot
	var err error
	if appoint.DoctorId != "" && appoint.PatientName != "" && appoint.Mobile != "" {
		existId := `SELECT doctor_name, spec_name, day, from_time, to_time, address from doctor_ast where doctor_id=$1;`
		row := db.QueryRow(existId, appoint.DoctorId)
		err = row.Scan(&sehedule.DoctorName, &sehedule.Specialization, &sehedule.Day, &sehedule.AvaliableFromTime, &sehedule.AvaliableToTime, &sehedule.Address)
		if err != nil {
			fmt.Println(err)
		}
		//strat next process
		//fmt.Println(sehedule.AvaliableFromTime, sehedule.AvaliableToTime)
		slotExist := `SELECT book_slot FROM book_ast where doctor_id=$1;`
		//fmt.Println(")))))", appoint.DoctorId)
		row = db.QueryRow(slotExist, appoint.DoctorId)
		err = row.Scan(&book.BookSlot)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("&&&&", book.BookSlot)
		if book.BookSlot == "" {
			InsertBookSlot := `INSERT INTO book_ast(doctor_id, book_slot) VALUES ($1, $2)`
			tTime := sehedule.AvaliableFromTime + ":" + "15"
			_, err := db.Exec(InsertBookSlot, appoint.DoctorId, tTime)
			if err != nil {
				fmt.Println(err)
			}
			book = BookSlot{
				DoctorId: appoint.DoctorId,
				BookSlot: tTime}
		}
		sehedule = DoctorSehedule{
			DoctorName:        sehedule.DoctorName,
			Specialization:    sehedule.Specialization,
			Day:               sehedule.Day,
			AvaliableFromTime: sehedule.AvaliableFromTime,
			AvaliableToTime:   sehedule.AvaliableToTime,
			Address:           sehedule.Address}

	}
	res := DoctorApponimntBooked{
		DoctorDetail: sehedule,
		SlotBooked:   book}
	return &res, nil
}
