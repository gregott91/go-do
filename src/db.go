package godo

import (
	"encoding/binary"
	"encoding/json"

	"github.com/boltdb/bolt"
)

const (
	dbName         = "godo.db"
	noteBucketName = "Notes"
)

// OpenDB creates and returns the DB
func OpenDB() (*bolt.DB, error) {
	dbPath, err := ConcatenateFileWithCurrentExeDir(dbName)
	if err != nil {
		return &bolt.DB{}, err
	}

	return bolt.Open(dbPath, 0600, nil)
}

// CreateBuckets create the mandatory buckets to be used by the DB
func CreateBuckets(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		var err error
		if err = createBucket(tx, noteBucketName); err != nil {
			return err
		}

		return nil
	})
}

func createBucket(tx *bolt.Tx, bucketName string) error {
	_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
	return err
}

// GetNotes gets all notes from the DB
func GetNotes(db *bolt.DB) ([]*Note, error) {
	var notes []*Note
	err := db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(noteBucketName))

		return b.ForEach(func(k, v []byte) error {
			var note *Note
			if err := json.Unmarshal(v, &note); err != nil {
				return err
			}

			notes = append(notes, note)
			return nil
		})
	})

	return notes, err
}

// InsertNote creates a note in the DB
func InsertNote(db *bolt.DB, note *Note) (int, error) {
	var insertID int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(noteBucketName))

		id, _ := b.NextSequence()
		note.ID = int(id)
		insertID = note.ID

		// Marshal user data into bytes.
		buf, err := json.Marshal(note)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put(itob(note.ID), buf)
	})

	return insertID, err
}

// DeleteNote deletes a note from the DB
func DeleteNote(db *bolt.DB, id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(noteBucketName))

		return b.Delete(itob(id))
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
