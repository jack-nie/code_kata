package models

type Tag struct {
	BaseModel
	Name       string
	PostsCount uint32 `gorm:"not null;default 0"`
}
