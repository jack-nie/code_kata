package models

type Post struct {
	BaseModel
	Title         string
	Content       string
	CommentsCount uint32 `gorm:"not null; default: 0"`
	Comments      []Comment
	CategoryID    uint
	Category      Category
	AdminID       uint
	Admin         Admin
	Tags          []Tag `gorm:"many2many:post_tags;"`
}
