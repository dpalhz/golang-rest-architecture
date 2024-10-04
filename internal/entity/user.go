package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"type:varchar(100)"`
	Username  string         `gorm:"type:varchar(100);unique"`
	Email     string         `gorm:"type:varchar(100);unique"`
	Password  string         `gorm:"type:varchar(100)"`
	Photo     string        `gorm:"type:varchar(255);null"` 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
