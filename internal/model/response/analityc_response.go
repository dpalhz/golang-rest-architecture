package response

import "time"

type Analytic struct {
	ID        uint      `json:"id"`
	BlogID    uint      `json:"blog_id"`
	UserID    uint      `json:"user_id"`
	DeviceType string    `json:"device_type"`
	Browser    string    `json:"browser"`
	IP         string    `json:"ip"`
	Referrer   string   `json:"referrer"`
	CreatedAt  time.Time `json:"created_at"`
}