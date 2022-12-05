package model

// EmployeeEventQuestion model for employee_event_questions table
type EmployeeEventQuestion struct {
	BaseModel

	EmployeeEventReviewerID UUID
	Content                 string
	Answer                  string
	Note                    string
}
