package entity

type Admin struct {
	User
	Blogs []Blog `gorm:"foreignKey:AdminID"`
}
