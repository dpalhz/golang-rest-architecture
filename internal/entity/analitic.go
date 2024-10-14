package entity

import "time"

type Analytic struct {
	ID        uint      `gorm:"primaryKey"`
	BlogID    uint      
	UserID    uint       
	DeviceType string    
	Browser    string    
	IP         string    
	Referrer   string   
	CreatedAt  time.Time 
}
