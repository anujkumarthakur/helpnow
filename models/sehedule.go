package models

import (
	"helpnow/config"
	"log"
)

type DoctorSehedule struct {
	DoctorName        string `json:"doctor_name"`
	Specialization    string `json:"spec_name"`
	Day               string `json:"day"`
	AvaliableFromTime string `json:"from_time"`
	AvaliableToTime   string `json:"to_time"`
	Address           string `json:"address"`
}

func SeheduleDoctorTime(DST DoctorSehedule) (*DoctorSehedule, error) {
	if DST.DoctorName != "" && DST.Specialization != "" && DST.Day != "" && DST.AvaliableFromTime != "" && DST.AvaliableToTime != "" && DST.Address != "" {
		db := config.GetDB()
		doctorQuery := `INSERT INTO doctor_ast (doctor_name, spec_name, day, from_time, to_time, address)
						VALUES ($1, $2, $3, $4, $5, $6)`
		_, err := db.Exec(doctorQuery, DST.DoctorName, DST.Specialization, DST.Day, DST.AvaliableFromTime, DST.AvaliableToTime, DST.Address)
		if err != nil {
			log.Println(err)
			// panic(err)
		}
	}
	return &DST, nil
}

// create table doctor_ast(doctor_id SERIAL PRIMARY KEY,
// 	doctor_name VARCHAR(200),
// 	spec_name VARCHAR(200),
// 	day VARCHAR(200),
// 	from_time VARCHAR(20),
// 	to_time VARCHAR(20),
// 	address VARCHAR(200));

// 	create table book_ast(appoiment_id SERIAL PRIMARY KEY,
// 		doctor_id VARCHAR(200),
// 		book_slot VARCHAR(200));
