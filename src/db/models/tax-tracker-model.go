package models

import (
	"time"
)

type TaxTravelTracker struct {
	ID                     int       `gorm:"primaryKey;autoIncrement"`
	CreatedAt              time.Time `gorm:"type:timestamp;autoUpdateTime"`
	CreatedBy              string    `gorm:"size:200"`
	TravelTime             string 	 `gorm:"size:500"`
	TravelFrom             string    `gorm:"size:500"`
	TravelTo               string    `gorm:"size:500"`
	Comment                string    `gorm:"size:500"`
	TravelDistance         float64
	EstimatedTaxDeductions float64
}
