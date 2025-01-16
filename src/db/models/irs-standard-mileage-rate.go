package models

type IrsStandardMileageRate struct {
	Year         int `gorm:"column:YEAR;primaryKey"`
	CentsPerMile int `gorm:"column:CENTS_PER_MILE"`
}


func (IrsStandardMileageRate) TableName() string {
    return "IRS_STANDARD_MILEAGE_RATE"
}