package repository

import (
	"tax-travel-tracker/src/db"
	"tax-travel-tracker/src/db/models"
	"time"
)

func CreateTravelRecord(record *models.TaxTravelTracker) error {
	// Ensure ID and CreatedAt are not set by the user
	record.ID = 0 // Let the database handle ID auto-increment
	record.CreatedAt = time.Time{} // Ensure CreatedAt is set by the database

	// Save the record to the database
	return db.DB.Create(record).Error
}

func GetTravelRecordByID(id int) (*models.TaxTravelTracker, error) {
	var record models.TaxTravelTracker
	if err := db.DB.First(&record, id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}
func GetAllTravelRecords() ([]models.TaxTravelTracker, error) {
	var records []models.TaxTravelTracker
	if err := db.DB.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func UpdateTravelRecord(record *models.TaxTravelTracker) error {
	// Create a map of fields that are allowed to be updated
	allowedUpdates := map[string]interface{}{
		"travel_time":             record.TravelTime,
		"travel_from":             record.TravelFrom,
		"travel_to":               record.TravelTo,
		"travel_distance":         record.TravelDistance,
		"estimated_tax_deductions": record.EstimatedTaxDeductions,
	}

	// Perform a scoped update to prevent modifications to restricted fields
	return db.DB.Model(&models.TaxTravelTracker{}).Where("id = ?", record.ID).Updates(allowedUpdates).Error
}

func DeleteTravelRecord(id int) error {
	return db.DB.Delete(&models.TaxTravelTracker{}, id).Error
}
