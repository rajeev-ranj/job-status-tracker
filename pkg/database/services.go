package database

import "fmt"

// createService creates a new service
func (db *DB) CreateService(service *Service) error {
	query := `INSERT INTO services (service_name, description, created_by) VALUES ($1, $2, $3) RETURNING service_id`
	err := db.QueryRow(query, service.ServiceName, service.Description, service.CreatedBy).Scan(&service.ServiceID)
	if err != nil {
		return fmt.Errorf("error creating service: %w", err)
	}
	return nil
}

// getService retrieves all services.
func (db *DB) GetServices() ([]*Service, error) {
	query := `SELECT service_id, service_name, description, created_by FROM services`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error retrieving services: %w", err)
	}

	defer rows.Close()

	var services []*Service
	for rows.Next() {
		var service Service
		if err := rows.Scan(&service.ServiceID, &service.ServiceName, &service.Description, &service.CreatedBy); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		services = append(services, &service)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error retrieving services: %w", err)
	}

	return services, nil
}

// getService retrieves a service by ID.
func (db *DB) GetService(serviceID int) (*Service, error) {
	query := `SELECT service_id, service_name, description, created_by FROM services WHERE service_id = $1`
	var service Service
	if err := db.QueryRow(query, serviceID).Scan(&service.ServiceID, &service.ServiceName, &service.Description, &service.CreatedBy); err != nil {
		return nil, fmt.Errorf("error retrieving service: %w", err)
	}
	return &service, nil
}

// updateService updates a service.
func (db *DB) UpdateService(service *Service) error {
	query := `UPDATE services SET service_name = $1, description = $2, created_by = $3 WHERE service_id = $4`
	_, err := db.Exec(query, service.ServiceName, service.Description, service.CreatedBy, service.ServiceID)
	if err != nil {
		return fmt.Errorf("error updating service: %w", err)
	}
	return nil
}

// deleteService deletes a service.
func (db *DB) DeleteService(serviceID int) error {
	query := `DELETE FROM services WHERE service_id = $1`
	_, err := db.Exec(query, serviceID)
	if err != nil {
		return fmt.Errorf("error deleting service: %w", err)
	}
	return nil
}
