package models

type Comment struct {
	BaseModel
	Content string
	UserID  uint
	PostID  uint
	Replys  []*Comment `gorm:"many2many:comment_replys;association_jointable_foreignkey:reply_id"`
}
