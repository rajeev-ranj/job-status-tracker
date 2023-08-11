package database

import (
	"database/sql"
	"time"
)

// Service represents a tracked service.
type Service struct {
	ServiceID   int    `json:"service_id"`
	ServiceName string `json:"service_name"`
	Description string `json:"description"`
	CreatedBy   int    `json:"created_by"`
}

// Job represents a tracked job.
type Job struct {
	JobID                  int          `json:"job_id"`
	ServiceID              int          `json:"service_id"`
	JobName                string       `json:"job_name"`
	ExpectedCompletionTime time.Time    `json:"expected_completion_time"`
	ActualCompletionTime   sql.NullTime `json:"actual_completion_time"`
	Status                 string       `json:"status"`
	Frequency              string       `json:"frequency"`
}

// User represents a user.
type User struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"` // Role of the user
}

// JobHistory represents a record in the job_history table.
type JobHistory struct {
	ID          int       `json:"id"`
	JobID       int       `json:"job_id"`
	Status      string    `json:"status"`
	Timestamp   time.Time `json:"timestamp"`
	Description string    `json:"description"`
	// any other relevant fields
}
