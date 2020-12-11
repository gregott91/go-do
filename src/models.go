package godo

import "time"

// Note represents an active note in the DB
type Note struct {
	ID        int
	Timestamp time.Time
	Text      string
}
