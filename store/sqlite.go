package store

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStore struct {
	db *sql.DB
}

func New(dbPath string) (Store, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("running migrations: %w", err)
	}

	return &SQLiteStore{db: db}, nil
}

func migrate(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		body TEXT NOT NULL DEFAULT ''
	);`
	_, err := db.Exec(query)
	return err
}

func (s *SQLiteStore) GetNotes() ([]Note, error) {
	rows, err := s.db.Query("SELECT id, title, body FROM notes ORDER BY id DESC")
	if err != nil {
		return nil, fmt.Errorf("querying notes: %w", err)
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Body); err != nil {
			return nil, fmt.Errorf("scanning note: %w", err)
		}
		notes = append(notes, n)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating notes: %w", err)
	}
	return notes, nil
}

func (s *SQLiteStore) SaveNote(note Note) error {
	query := `
	INSERT INTO notes (id, title, body) VALUES (?,?,?)
	ON CONFLICT(id) DO UPDATE SET
		title = excluded.title,
		body = excluded.body`

	_, err := s.db.Exec(query, note.ID, note.Title, note.Body)
	if err != nil {
		return fmt.Errorf("saving note: %w", err)
	}
	return nil
}
