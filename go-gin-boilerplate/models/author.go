package models

type Author struct {
	BaseModel
	Name       string `gorm:"not null"`
	BirthDay   string
	Gender     string
	PhotoURL   string
	Active     bool
	Password   string
	Posts      []Post
	PostsCount uint32 `gorm:"not null;default 0;"`
}
