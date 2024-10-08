package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"type:varchar(100);unique"` 
	UpdatedAt time.Time
}

type Tag struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"type:varchar(100);unique"` 
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Blog struct {
	ID         uint           `gorm:"primaryKey"`
	Title      string         `gorm:"type:varchar(255)"`
	Content    string         `gorm:"type:text"`  
	ReadTime   int            `gorm:"default:0"`              
	IsBlocked  bool           `gorm:"default:false"`          
	AdminID    uint                                           
	Categories []Category     `gorm:"many2many:blog_categories;"` 
	Tags       []Tag          `gorm:"many2many:blog_tags;"`       
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}