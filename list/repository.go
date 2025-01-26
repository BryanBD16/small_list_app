package list

//Command to create a database
//mysql -u root -p -e 'CREATE DATABASE name_of_db'
//Command to run query
//mysql -u user_name -p password -e 'SQL Query' database
//mysql -u root -p -e 'CREATE TABLE elements(id INT AUTO_INCREMENT PRIMARY KEY,name VARCHAR(255) NOT NULL, description VARCHAR(255));' list_app

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type IRepository interface {
	AddElement(elem Element) error
	GetElements() ([]Element, error)
	ClearList() error
}

type Repository struct {
	db *sql.DB
}

// NewRepository initializes the database connection.
func NewRepository(dsn string) (*Repository, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Repository{db: db}, nil
}

// AddElement inserts a new element into the database.
func (r *Repository) AddElement(elem Element) error {
	_, err := r.db.Exec("INSERT INTO elements (name, description) VALUES (?, ?)", elem.Name, elem.Description)
	if err != nil {
		return fmt.Errorf("failed to insert element: %w", err)
	}
	return nil
}

// GetElements retrieves all elements from the database.
func (r *Repository) GetElements() ([]Element, error) {
	rows, err := r.db.Query("SELECT name, description FROM elements")
	if err != nil {
		return nil, fmt.Errorf("failed to query elements: %w", err)
	}
	defer rows.Close()

	var elements []Element
	for rows.Next() {
		var elem Element
		// Scan both name and description into the Element struct
		if err := rows.Scan(&elem.Name, &elem.Description); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		elements = append(elements, elem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return elements, nil
}

// ClearTable deletes all rows from the elements table.
func (r *Repository) ClearList() error {
	_, err := r.db.Exec("DELETE FROM elements")
	if err != nil {
		return fmt.Errorf("failed to clear elements table: %w", err)
	}
	return nil
}

// Close closes the database connection.
func (r *Repository) Close() error {
	return r.db.Close()
}
