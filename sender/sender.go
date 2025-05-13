package sender

// CreateAndScheduleSummaryAppointmentMessages
// Creates and schedules summary appointment messages, and outputs them to tables. The messages are for appointments on a given date (passed as a parameter to your binary)
// Expected output format:
//
// MESSAGE FOR A DOCTOR FOR A GIVEN CENTER (THE APPOINTMENTS SHOULD BE SORTED BY TIME):
// Dr. <DoctorName>'s appointments on <Date> at <CenterName>: <Number of appointments>
// <Time>, <Duration>: <PatientName> (<TreatmentCategory>)
// <Time>, <Duration>: <PatientName> (<TreatmentCategory>)
// ...so on.
//
// Example:
// Dr. John's appointments on <12th May, 2025> at Aundh: 3
// 10:00 am, 1h 30m: Ms. Alice (Consultation)
// 11:30 am, 15m: Mr. Bob
// 12:00 pm, 1h: Mr. Charlie (Ortho)
//
// Notice that Mr. Bob's appointment does not have a category in the message. This is because the treatment category for this appointment was "Not Specified". If an appointment has this treatment category, you must leave the category blank in the message.
//
// SUMMARY OF THE DAY FOR A GIVEN LOCATION:
// Summary of appointments at <CenterName> on <Date>: <Number of appointments>
// Dr. <DoctorName>: <Number of appointments>
// Dr. <DoctorName>: <Number of appointments>
//
// Example:
// Summary of appointments at Aundh on 12th May, 2025: 3
// Dr. John: 3
//
// You can insert all messages sent to doctors in a table called DoctorMessages. It should have at least the person ID of the doctor and the doctor's phone number along with the message
// You can insert all messages that summarize a center into a table called CenterMessages. It should contain at least the center ID of the center along with the message
// TODO Change function input parameter type if needed
func CreateAndScheduleSummaryAppointmentMessages(appointments []any) (err error) {
	panic("implement me")
}
