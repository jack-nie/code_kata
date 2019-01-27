package models

type Category struct {
	BaseModel
	Name       string
	PostsCount uint32 `gorm:"not null;default: 0"`
}
