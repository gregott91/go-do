package godo

import (
	"time"

	"github.com/boltdb/bolt"
)

// Note represents an active note in the DB
type Note struct {
	ID        int
	Timestamp time.Time
	Text      string
}

// NotesConnection maintains a connection to the DB
type NotesConnection struct {
	DB *bolt.DB
}

// InitializeNotesBackend creates and returns the DB
func InitializeNotesBackend() (*NotesConnection, error) {
	db, err := OpenDB()

	if err != nil {
		return &NotesConnection{}, err
	}

	if err = CreateBuckets(db); err != nil {
		return &NotesConnection{}, err
	}

	return &NotesConnection{DB: db}, nil
}

// CreateNote creates a note in the DB
func (conn *NotesConnection) CreateNote(text string) (*Note, error) {
	note := &Note{
		Text:      text,
		Timestamp: time.Now(),
	}

	id, err := InsertNote(conn.DB, note)

	note.ID = id

	return note, err
}

// CloseConnection closes a connection to the DB
func (conn *NotesConnection) CloseConnection() {
	conn.DB.Close()
}

// GetNotes gets all notes from the DB
func (conn *NotesConnection) GetNotes() ([]*Note, error) {
	return GetNotes(conn.DB)
}

// RemoveNote removes a note from the DB
func (conn *NotesConnection) RemoveNote(id int) error {
	return DeleteNote(conn.DB, id)
}
