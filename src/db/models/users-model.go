package models

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"size:200;not null"`
	Password string `gorm:"size:200;not null"`
}

