package response

import "time"

type BlogDetail struct {
	ID         uint     `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	ReadTime   int      `json:"read_time"`
	IsBlocked  bool     `json:"is_blocked"`
	Categories []string `json:"categories"`
	Tags       []string `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
}

type Blog struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	ReadTime   int       `json:"read_time"`
	IsBlocked  bool      `json:"is_blocked"`
	AdminID    uint      `json:"admin_id"`
	Categories []string  `json:"categories"`
	Tags       []string  `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}




type Blogs struct {
	Blogs      []BlogDetail `json:"blogs"`      
	TotalCount int64         `json:"total_count"` 
	Page       int           `json:"page"`        
	PageSize   int           `json:"page_size"`   
}

