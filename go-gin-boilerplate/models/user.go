package models

import (
	"fmt"

	"github.com/jack-nie/code_kata/go-gin-boilerplate/forms"
	"github.com/sirupsen/logrus"
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

func (u User) Signup(userPayload forms.UserSignup) (*User, error) {
	user := &User{
		Name:     userPayload.Name,
		BirthDay: userPayload.BirthDay,
		Gender:   userPayload.Gender,
		PhotoURL: userPayload.PhotoURL,
		Active:   true,
	}
	err := db.Create(&user).Error
	if err != nil {
		logrus.Fatalf(fmt.Sprintf("server exception, %v", err))
		return nil, err
	}
	return user, nil
}
