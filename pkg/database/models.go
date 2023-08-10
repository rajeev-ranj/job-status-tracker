package database

import "time"

// Service represents a tracked service.
type Service struct {
	ServiceID   int    `json:"service_id"`
	ServiceName string `json:"service_name"`
	Description string `json:"description"`
	CreatedBy   int    `json:"created_by"`
}

// Job represents a tracked job.
type Job struct {
	JobID                  int       `json:"job_id"`
	ServiceID              int       `json:"service_id"`
	JobName                string    `json:"job_name"`
	ExpectedCompletionTime time.Time `json:"expected_completion_time"`
	ActualCompletionTime   time.Time `json:"actual_completion_time"`
	Status                 string    `json:"status"`
	Frequency              string    `json:"frequency"`
}

// User represents a user.
type User struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"` // Role of the user
}

// job_history represents a job history.
type JobHistory struct {
}
