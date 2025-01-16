package repository

import (
	"tax-travel-tracker/src/db"
	"tax-travel-tracker/src/db/models"
)

func GetMilageRecordByID(id int) (*models.IrsStandardMileageRate, error) {
	var record models.IrsStandardMileageRate
	if err := db.DB.First(&record, id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func GetAllMileageRecords() ([]models.IrsStandardMileageRate, error) {
	var records []models.IrsStandardMileageRate
	if err := db.DB.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}
