package models

import (
	"helpnow/config"
	"log"
)

func GetDoctorsList() ([]DoctorSehedule, error) {
	var err error
	var doctor DoctorSehedule
	var data []DoctorSehedule
	db := config.GetDB()
	doctorlistQuery := `SELECT doctor_name, spec_name, day, from_time, to_time, address FROM doctor_ast`

	rows, err := db.Query(doctorlistQuery)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&doctor.DoctorName, &doctor.Specialization, &doctor.Day, &doctor.AvaliableFromTime, &doctor.AvaliableToTime, &doctor.Address)
		if err != nil {
			log.Println(err)
		}
		res := DoctorSehedule{
			DoctorName:        doctor.DoctorName,
			Specialization:    doctor.Specialization,
			Day:               doctor.Day,
			AvaliableFromTime: doctor.AvaliableFromTime,
			AvaliableToTime:   doctor.AvaliableToTime,
			Address:           doctor.Address}

		data = append(data, res)
	}
	return data, nil

}
