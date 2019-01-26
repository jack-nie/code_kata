package models

import (
	"fmt"
	"time"

	"github.com/jack-nie/code_kata/go-gin-boilerplate/forms"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type User struct {
	ID        string `json:"user_id,omitempty"`
	Name      string `json:"name"`
	BirthDay  string `json:"birthday"`
	Gender    string `json:"gender"`
	PhotoURL  string `json:"photo_url"`
	Time      int64  `json:"current_time"`
	Active    bool   `json:"active,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (u User) Signup(userPayload forms.UserSignup) (*User, error) {
	id := uuid.NewV4()
	user := &User{
		ID:        id.String(),
		Name:      userPayload.Name,
		BirthDay:  userPayload.BirthDay,
		Gender:    userPayload.Gender,
		PhotoURL:  userPayload.PhotoURL,
		Time:      time.Now().UnixNano(),
		Active:    true,
		UpdatedAt: time.Now().UnixNano(),
	}
	err := db.Create(&user).Error
	if err != nil {
		logrus.Error(fmt.Sprintf("server exception, %v", err))
	}
	return user, nil
}
