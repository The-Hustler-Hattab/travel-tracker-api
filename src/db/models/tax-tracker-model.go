package models

import (
	"time"
)

type TaxTravelTracker struct {
	ID                    int       `gorm:"primaryKey;autoIncrement"`
	CreatedAt             time.Time `gorm:"type:timestamp;autoUpdateTime"`
	CreatedBy             string    `gorm:"size:200"`
	TravelTime            time.Time `gorm:"type:date"`
	TravelFrom            string    `gorm:"size:500"`
	TravelTo              string    `gorm:"size:500"`
	TravelDistance        int
	EstimatedTaxDeductions int
}
