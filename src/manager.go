package godo

import (
	"time"

	"github.com/boltdb/bolt"
)

// NotesConnection maintains a connection to the DB
type NotesConnection struct {
	CreateNote      func(text string) (*Note, error)
	GetNotes        func() ([]*Note, error)
	CloseConnection func()
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

	return &NotesConnection{
		CreateNote: func(text string) (*Note, error) {
			return CreateNote(db, text)
		},
		CloseConnection: func() {
			db.Close()
		},
		GetNotes: func() ([]*Note, error) {
			return GetNotes(db)
		},
	}, nil
}

// CreateNote creates a note in the DB
func CreateNote(db *bolt.DB, text string) (*Note, error) {
	note := &Note{
		Text:      text,
		Timestamp: time.Now(),
	}

	id, err := InsertNote(db, note)

	note.ID = id

	return note, err
}
