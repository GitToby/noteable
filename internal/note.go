package internal

import (
	"time"

	"github.com/spf13/viper"

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
	// search for all notes containing the search term
	db.Where("content LIKE ?", "%"+term+"%").Limit(limit).Find(&notes)
	return notes
}

func getDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(viper.GetString("database_path")), &gorm.Config{})
	db.AutoMigrate(&Note{})

	if err == nil {
		return db
	}
	panic("no connection to db")
}
