package request

type Blogs struct {
	Page  int `query:"page" validate:"required,min=1"`
	Limit int `query:"limit" validate:"required,min=1"`
}

type CreateBlog struct {
	Title      string `json:"title" validate:"required,max=255"`
	Content    string `json:"content" validate:"required"`
	ReadTime   int    `json:"read_time" validate:"min=0"`
	IsBlocked  bool   `json:"is_blocked"`
	AdminID    uint   `json:"admin_id" validate:"required"`
	Categories []uint `json:"categories" validate:"required"`
	Tags       []uint `json:"tags" validate:"required"`
}

type UpdateBlog struct {
	Title     string `json:"title" validate:"required,max=255"`
	Content   string `json:"content" validate:"required"`
	ReadTime  int    `json:"read_time" validate:"min=0"`
	IsBlocked bool   `json:"is_blocked"`
}

type BlogFilter struct {
	Categories []uint `json:"categories" validate:"required"`
	Tags       []uint `json:"tags" validate:"required"`
	MatchAll   bool   `json:"match_all"`
}
