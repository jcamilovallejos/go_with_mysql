package controllers

import (
	"database/sql"
	"fmt"
	"go_with_mysql/models"
)

func ListContacts(db *sql.DB) ([]models.Contact, error) {
	query := "SELECT id, name, email, phone FROM contact"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var contacts []models.Contact

	for rows.Next() {
		var contact models.Contact
		if err := rows.Scan(
			&contact.Id,
			&contact.Name,
			&contact.Email,
			&contact.Phone,
		); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		contacts = append(contacts, contact)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return contacts, nil
}

func GetContactByID(db *sql.DB, id int64) (models.Contact, error) {
	query := "SELECT id, name, email, phone FROM contact WHERE id = ?"

	var contact models.Contact
	if err := db.QueryRow(query, id).Scan(
		&contact.Id,
		&contact.Name,
		&contact.Email,
		&contact.Phone,
	); err != nil {
		if err == sql.ErrNoRows {
			return models.Contact{}, fmt.Errorf("contact not found with ID %d", id)
		}
		return models.Contact{}, fmt.Errorf("error retrieving contact: %w", err)
	}

	return contact, nil
}

func CreateContact(db *sql.DB, contact models.Contact) (int64, error) {
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"

	result, err := db.Exec(
		query,
		contact.Name,
		contact.Email,
		contact.Phone,
	)
	if err != nil {
		return 0, fmt.Errorf("error executing insert: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert ID: %w", err)
	}

	return id, nil
}

func UpdateContact(db *sql.DB, contact models.Contact) (models.Contact, error) {
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	result, err := db.Exec(
		query,
		contact.Name,
		contact.Email,
		contact.Phone,
		contact.Id,
	)
	if err != nil {
		return models.Contact{}, fmt.Errorf("error executing update: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Contact{}, fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return models.Contact{}, fmt.Errorf("contact not found with ID %d", contact.Id)
	}

	return contact, nil
}

func DeleteContact(db *sql.DB, contactID int64) (int64, error) {
	query := "DELETE FROM contact WHERE id = ?"

	result, err := db.Exec(query, contactID)
	if err != nil {
		return 0, fmt.Errorf("error executing delete: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("contact not found with ID %d", contactID)
	}

	return contactID, nil
}
