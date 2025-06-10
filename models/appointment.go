package models

import "time"

type Appointment struct {
	AppointmentID     int
	CenterID          int
	CenterName        string
	DoctorID          int
	DoctorName        string
	DoctorMobile      string
	PatientName       string
	StartTime         time.Time
	EndTime           time.Time
	Status            string
	TreatmentCategory string
}

type DoctorMessage struct {
	DoctorID    int
	DoctorPhone string
	Message     string
}

type CenterMessage struct {
	CenterID int
	Message  string
}
