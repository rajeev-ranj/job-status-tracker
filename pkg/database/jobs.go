package database

import (
	"fmt"
)

// CreateJob creates a new job.
func (db *DB) CreateJob(job *Job) error {
	query := `INSERT INTO jobs (service_id, job_name, expected_completion_time, actual_completion_time, status, frequency) VALUES ($1, $2, $3, $4, $5, $6) RETURNING job_id`
	err := db.QueryRow(query, job.ServiceID, job.JobName, job.ExpectedCompletionTime, job.ActualCompletionTime, job.Status, job.Frequency).Scan(&job.JobID)
	if err != nil {
		return fmt.Errorf("error creating job: %w", err)
	}
	return nil
}

// GetJobs retrieves all jobs.
func (db *DB) GetJobs() ([]*Job, error) {
	query := `SELECT job_id, service_id, job_name, expected_completion_time, actual_completion_time, status, frequency FROM jobs`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error retrieving jobs: %w", err)
	}
	defer rows.Close()

	var jobs []*Job
	for rows.Next() {
		var job Job
		if err := rows.Scan(&job.JobID, &job.ServiceID, &job.JobName, &job.ExpectedCompletionTime, &job.ActualCompletionTime, &job.Status, &job.Frequency); err != nil {
			return nil, fmt.Errorf("error scanning job: %w", err)
		}
		jobs = append(jobs, &job)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating jobs: %w", err)
	}

	return jobs, nil
}

// GetJob retrieves a job by ID.
func (db *DB) GetJob(jobID int) (*Job, error) {
	query := `SELECT job_id, service_id, job_name, expected_completion_time, actual_completion_time, status, frequency FROM jobs WHERE job_id = $1`
	var job Job
	if err := db.QueryRow(query, jobID).Scan(&job.JobID, &job.ServiceID, &job.JobName, &job.ExpectedCompletionTime, &job.ActualCompletionTime, &job.Status, &job.Frequency); err != nil {
		return nil, fmt.Errorf("error retrieving job: %w", err)
	}
	return &job, nil
}

// UpdateJob updates a job.
func (db *DB) UpdateJob(job *Job) error {
	query := `UPDATE jobs SET service_id = $1, job_name = $2, expected_completion_time = $3, actual_completion_time = $4, status = $5, frequency = $6 WHERE job_id = $7`
	_, err := db.Exec(query, job.ServiceID, job.JobName, job.ExpectedCompletionTime, job.ActualCompletionTime, job.Status, job.Frequency, job.JobID)
	if err != nil {
		return fmt.Errorf("error updating job: %w", err)
	}
	return nil
}

// DeleteJob deletes a job.
func (db *DB) DeleteJob(jobID int) error {
	query := `DELETE FROM jobs WHERE job_id = $1`
	_, err := db.Exec(query, jobID)
	if err != nil {
		return fmt.Errorf("error deleting job: %w", err)
	}
	return nil
}

// GetJobsByServiceID retrieves all jobs for a service
func (db *DB) GetJobsByServiceID(serviceID int) ([]*Job, error) {
	query := `SELECT job_id, service_id, job_name, expected_completion_time, actual_completion_time, status, frequency FROM jobs WHERE service_id = $1`
	rows, err := db.Query(query, serviceID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving jobs: %w", err)
	}
	defer rows.Close()

	var jobs []*Job
	for rows.Next() {
		var job Job
		if err := rows.Scan(&job.JobID, &job.ServiceID, &job.JobName, &job.ExpectedCompletionTime, &job.ActualCompletionTime, &job.Status, &job.Frequency); err != nil {
			return nil, fmt.Errorf("error scanning job: %w", err)
		}
		jobs = append(jobs, &job)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating jobs: %w", err)
	}

	return jobs, nil
}

// GetJobHistory retrieves the history of a specific job by its ID.
func (db *DB) GetJobHistory(jobID int) ([]JobHistory, error) {
	rows, err := db.Query(`SELECT id, job_id, status, timestamp, description FROM job_history WHERE job_id = $1 ORDER BY timestamp DESC`, jobID)
	if err != nil {
		return nil, fmt.Errorf("error querying job history: %w", err)
	}
	defer rows.Close()

	var history []JobHistory
	for rows.Next() {
		var record JobHistory
		if err := rows.Scan(&record.ID, &record.JobID, &record.Status, &record.Timestamp, &record.Description); err != nil {
			return nil, fmt.Errorf("error scanning job history row: %w", err)
		}
		history = append(history, record)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through job history rows: %w", err)
	}

	return history, nil
}
