package models

import (
	"github.com/jack-nie/code_kata/go-gin-boilerplate/forms"
)

type User struct {
	BaseModel
	Name     string `gorm:"not null"`
	BirthDay string
	Gender   string
	PhotoURL string
	Active   bool
	Password string
	Comments []Comment
}

func (u User) Signup(userPayload forms.UserSignup) (user *User, err error) {
	user = &User{
		Name:     userPayload.Name,
		BirthDay: userPayload.BirthDay,
		Gender:   userPayload.Gender,
		PhotoURL: userPayload.PhotoURL,
		Active:   true,
	}
	if err = db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
