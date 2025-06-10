package main

import (
	"AppointmentSummary_Assignment/database"
	"AppointmentSummary_Assignment/sender"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <yyyy-mm-dd>")
		os.Exit(1)
	}
	date := os.Args[1]

	database.Init()

	appointments, err := database.ReadDataForDate(date)
	if err != nil {
		fmt.Printf("Error reading data: %v\n", err)
		os.Exit(1)
	}

	if len(appointments) == 0 {
		fmt.Println("No appointments found for the given date.")
		return
	}

	err = sender.CreateAndScheduleSummaryAppointmentMessages(date, appointments)
	if err != nil {
		fmt.Printf("Error generating summary messages: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Summary messages generated and stored successfully.")
}
