package database

import (
	"fmt"
)

// CreateUser creates a new user.
func (db *DB) CreateUser(user *User) error {
	query := `INSERT INTO users (username, email, role) VALUES ($1, $2, $3) RETURNING user_id`
	err := db.QueryRow(query, user.Username, user.Email, user.Role).Scan(&user.UserID)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

// GetUsers retrieves all users.
func (db *DB) GetUsers() ([]*User, error) {
	query := `SELECT user_id, username, email, role FROM users`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error retrieving users: %w", err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Role); err != nil {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating users: %w", err)
	}

	return users, nil
}

// GetUser retrieves a user by ID.
func (db *DB) GetUser(userID int) (*User, error) {
	query := `SELECT user_id, username, email, role FROM users WHERE user_id = $1`
	var user User
	if err := db.QueryRow(query, userID).Scan(&user.UserID, &user.Username, &user.Email, &user.Role); err != nil {
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}
	return &user, nil
}

// UpdateUser updates a user.
func (db *DB) UpdateUser(user *User) error {
	query := `UPDATE users SET username = $1, email = $2, role = $3 WHERE user_id = $4`
	_, err := db.Exec(query, user.Username, user.Email, user.Role, user.UserID)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

// DeleteUser deletes a user.
func (db *DB) DeleteUser(userID int) error {
	query := `DELETE FROM users WHERE user_id = $1`
	_, err := db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}
