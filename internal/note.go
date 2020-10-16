package internal

import (
	"time"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Note structure to pass around and save
type Note struct {
	gorm.Model
	Content string
}

// Save the given note to the database
func (n *Note) Save() {
	db := getDB()
	db.Create(n)
}

// GetNotesSince retrieves since a given time
func GetNotesSince(date time.Time) []Note {
	db := getDB()
	notes := []Note{}
	db.Where("created_at >= ?", date).Find(&notes)
	return notes
}

// SearchNotes searches the body of all notes
func SearchNotes(term string, limit int) []Note {
	db := getDB()
	notes := []Note{}
	db.Where("content LIKE ?", "%"+term+"%").Limit(limit).Find(&notes)
	return notes
}

func getDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Note{})

	if err == nil {
		return db
	}
	panic("no connection to db")
}
