Appointment Summary Generator - BestoSys Assignment
This Go project generates summary messages for doctors and centers based on appointment data for a given date. The summaries are then stored in the database.

🛠 Tech Stack
Language: Go (Golang)
Database: MySQL
Dependencies:
github.com/go-sql-driver/mysql
📁 Project Structure
⚙️ Setup Instructions
Clone the repository:
git clone https://github.com/JayeshJadhav1107/AppointmentSummary_Assignment.git
cd AppointmentSummary_Assignment
Run the project:
go run main.go 2025-05-12 (sample date) Replace 2025-05-12 with any valid date in YYYY-MM-DD format.

📌 Features Fetches appointment data for a given date. Groups data by Doctor and Center. Calculates total appointments, duration, and patient info. Stores summary messages in the database. Logs inserted messages to the console.

🧪 Example Output '1', '117800117', '8900000117', 'Dr. Beb Mem's appointments on 12 May, 2025 at Kharadi: 1\n9:30 pm, 1h 15m: Mr Qoh Tow (Oral Surgery)'

🙋 Author Jayesh Jadhav