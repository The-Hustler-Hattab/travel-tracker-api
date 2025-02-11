package repository

import (
	"tax-travel-tracker/src/db"
	"tax-travel-tracker/src/db/models"
)

func CreateUserRecord(record *models.User) error {
	// Ensure ID and CreatedAt are not set by the user
	record.ID = 0 // Let the database handle ID auto-increment

	// Save the record to the database
	return db.DB.Create(record).Error
}