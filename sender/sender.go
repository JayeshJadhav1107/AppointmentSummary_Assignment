package sender

import (
	"AppointmentSummary_Assignment/models"
	"database/sql"
	"fmt"
	"sort"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateAndScheduleSummaryAppointmentMessages(date string, appointments []models.Appointment) error {
	doctorMap := make(map[string][]models.Appointment)
	centerMap := make(map[int]map[int]int)
	centerTotal := make(map[int]int)
	centerNames := make(map[int]string)

	for _, appt := range appointments {
		key := fmt.Sprintf("%d_%d", appt.CenterID, appt.DoctorID)
		doctorMap[key] = append(doctorMap[key], appt)

		if centerMap[appt.CenterID] == nil {
			centerMap[appt.CenterID] = make(map[int]int)
		}
		centerMap[appt.CenterID][appt.DoctorID]++
		centerTotal[appt.CenterID]++
		centerNames[appt.CenterID] = appt.CenterName
	}

	var doctorMsgs []models.DoctorMessage
	var centerMsgs []models.CenterMessage

	parsedDate, _ := time.Parse("2006-01-02", date)
	dateStr := parsedDate.Format("2 Jan, 2006")

	for _, appts := range doctorMap {
		if len(appts) == 0 {
			continue
		}
		first := appts[0]
		header := fmt.Sprintf("Dr. %s's appointments on %s at %s: %d",
			first.DoctorName, dateStr, first.CenterName, len(appts),
		)

		sort.Slice(appts, func(i, j int) bool {
			return appts[i].StartTime.Before(appts[j].StartTime)
		})

		lines := []string{header}
		for _, a := range appts {
			start := a.StartTime.Format("3:04 pm")
			duration := a.EndTime.Sub(a.StartTime)
			h := int(duration.Hours())
			m := int(duration.Minutes()) % 60

			durStr := ""
			if h > 0 {
				durStr += fmt.Sprintf("%dh ", h)
			}
			if m > 0 {
				durStr += fmt.Sprintf("%dm", m)
			}
			durStr = strings.TrimSpace(durStr)

			category := ""
			if a.TreatmentCategory != "Not Specified" {
				category = fmt.Sprintf(" (%s)", a.TreatmentCategory)
			}

			line := fmt.Sprintf("%s, %s: %s%s", start, durStr, a.PatientName, category)
			lines = append(lines, line)
		}

		msg := strings.Join(lines, "\n")
		doctorMsgs = append(doctorMsgs, models.DoctorMessage{
			DoctorID:    first.DoctorID,
			DoctorPhone: first.DoctorMobile,
			Message:     msg,
		})
	}

	for centerID, docMap := range centerMap {
		centerName := centerNames[centerID]
		header := fmt.Sprintf("Summary of appointments at %s on %s: %d",
			centerName, dateStr, centerTotal[centerID],
		)

		lines := []string{header}
		for docID, count := range docMap {
			var doctorName string
			for _, a := range appointments {
				if a.CenterID == centerID && a.DoctorID == docID {
					doctorName = a.DoctorName
					break
				}
			}
			lines = append(lines, fmt.Sprintf("Dr. %s: %d", doctorName, count))
		}

		msg := strings.Join(lines, "\n")
		centerMsgs = append(centerMsgs, models.CenterMessage{
			CenterID: centerID,
			Message:  msg,
		})
	}

	return saveMessages(doctorMsgs, centerMsgs)
}

func saveMessages(dMsgs []models.DoctorMessage, cMsgs []models.CenterMessage) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bestosys")
	if err != nil {
		return err
	}
	defer db.Close()

	for _, msg := range dMsgs {
		fmt.Println("Inserting doctor message:", msg)
		_, err := db.Exec("INSERT INTO doctormessages (DoctorID, Mobile, Message) VALUES (?, ?, ?)",
			msg.DoctorID, msg.DoctorPhone, msg.Message)
		if err != nil {
			return err
		}
	}

	for _, msg := range cMsgs {
		fmt.Println("Inserting center message:", msg)
		_, err := db.Exec("INSERT INTO centermessages (CenterID, Message) VALUES (?, ?)",
			msg.CenterID, msg.Message)
		if err != nil {
			return err
		}
	}

	return nil
}
