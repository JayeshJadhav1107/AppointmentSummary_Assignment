package database

import (
	"AppointmentSummary_Assignment/models"
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/bestosys"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
}

func ReadDataForDate(date string) ([]models.Appointment, error) {
	query := `
		SELECT 
			a.AppointmentID, a.CenterID, c.Name AS CenterName,
			d.DoctorStaffID, d.Name AS DoctorName, d.Mobile AS DoctorMobile,
			CONCAT(p.Salutation, ' ', p.Name) AS PatientName,
			a.StartTime, a.EndTime, a.Status, a.TreatmentCategory
		FROM appointment a
		JOIN center c ON a.CenterID = c.CenterID
		JOIN doctorstaff d ON a.DoctorStaffID = d.DoctorStaffID
		JOIN patient p ON a.PatientID = p.PatientID
		WHERE DATE(a.StartTime) = ? AND a.Status != 'C'
	`

	rows, err := db.Query(query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []models.Appointment
	// layout := "2006-01-02 15:04:05"

	for rows.Next() {
		var a models.Appointment
		var startStr, endStr string

		err := rows.Scan(&a.AppointmentID, &a.CenterID, &a.CenterName,
			&a.DoctorID, &a.DoctorName, &a.DoctorMobile,
			&a.PatientName, &startStr, &endStr, &a.Status, &a.TreatmentCategory)
		if err != nil {
			return nil, err
		}

		a.StartTime, err = ParseFlexibleTime(startStr)
		if err != nil {
			log.Printf("Error parsing StartTime: %v", err)
			continue
		}

		a.EndTime, err = ParseFlexibleTime(endStr)
		if err != nil {
			log.Printf("Error parsing EndTime: %v", err)
			continue
		}

		appointments = append(appointments, a)
	}

	return appointments, nil
}
func ParseFlexibleTime(raw string) (time.Time, error) {
	formats := []string{
		"2006-01-02 15:04:05", // MySQL DATETIME
		"1/2/2006 15:04",      // e.g., 5/13/2025 10:00
		"01/02/2006 15:04",    // e.g., 05/13/2025 10:00
		"2006-1-2 15:04",      // e.g., 2025-5-13 10:00
	}

	for _, format := range formats {
		if t, err := time.Parse(format, raw); err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("unsupported time format: " + raw)
}
