package model

type APIResponse struct {
	ErrorCode int         `json:"status_code"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
}
