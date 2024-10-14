package request

type CreateAnalytic struct {
	BlogID     uint   `json:"blog_id" validate:"required"`
	UserID     uint   `json:"user_id" validate:"required"`
	DeviceType string `json:"device_type" validate:"required"`
	Browser    string `json:"browser" validate:"required"`
	IP         string `json:"ip" validate:"required"`
	Referrer   string `json:"referrer" validate:"required"`
}

type UpdateAnalytic struct {
	BlogID     uint   `json:"blog_id,omitempty"`
	UserID     uint   `json:"user_id,omitempty"`
	DeviceType string `json:"device_type,omitempty"`
	Browser    string `json:"browser,omitempty"`
	IP         string `json:"ip,omitempty"`
	Referrer   string `json:"referrer,omitempty"`
}