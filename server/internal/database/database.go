package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed migrations/000createtodo.sql
var migrationString string

// ErrIDNotFound is returned when the id is not found in the database
var ErrIDNotFound = errors.New("not found")

// TodoItem is the struct that represents a todo item
type TodoItem struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Finished int    `json:"finished"`
}

// Database is the struct that represents the database
type Database struct {
	db *sql.DB
}

// NewDB creates a new database
func NewDB(file string) (*Database, error) {
	db, err := sql.Open("sqlite3", file)

	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

// CloseDB closes the database
func (d *Database) CloseDB() {
	d.db.Close()
}

// Migrate migrates the database
func (d *Database) Migrate() error {

	_, err := d.db.Exec(migrationString)
	return err
}

// Insert inserts a new item in the database
func (d *Database) Insert(item TodoItem) (int64, error) {
	insertQuery := fmt.Sprintf(`INSERT INTO todo (name) VALUES ("%s")`, item.Name)

	res, err := d.db.Exec(insertQuery)

	if err != nil {
		return -1, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

// FindAll returns all the items in the database
func (d *Database) FindAll() ([]TodoItem, error) {
	rows, err := d.db.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todoList []TodoItem
	for rows.Next() {

		var item TodoItem
		err := rows.Scan(&item.ID, &item.Name, &item.Finished)

		if err != nil {
			return nil, err
		}
		todoList = append(todoList, item)

	}

	return todoList, nil
}


// FindById returns the item with the given id
func (d *Database) FindById(id int64) (*TodoItem, error) {
	findQuery := fmt.Sprintf("SELECT * FROM todo WHERE id=%d", id)

	// it doesn't return errors. errors are deffered until rows scan
	row := d.db.QueryRow(findQuery)

	var item TodoItem

	err := row.Scan(&item.ID, &item.Name, &item.Finished)

	if err != nil {
		return nil, fmt.Errorf("id %d %w", id, ErrIDNotFound)
	}
	return &item, nil
}


// Update updates the item with the given id
func (d *Database) Update(item TodoItem) error {

	updateQuery := fmt.Sprintf(`UPDATE todo SET name="%s", finished=%d WHERE id=%d`, item.Name, item.Finished, item.ID)

	res, err := d.db.Exec(updateQuery)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrIDNotFound
	}

	return err
}

// Delete deletes the item with the given id
func (d *Database) Delete(id string) error {
	deleteQuery := fmt.Sprintf(`DELETE FROM todo WHERE id=%s`, id)

	res, err := d.db.Exec(deleteQuery)

	if err != nil {
		return err
	}

	rowsEffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsEffected == 0 {
		return ErrIDNotFound
	}

	return err
}
