package models

import (
	"errors"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"

	"github.com/jack-nie/code_kata/go-gin-boilerplate/config"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/forms"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/utils"
)

type Admin struct {
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

type tokenClaims struct {
	AdminID string `json:"admin_id"`
	jwt.StandardClaims
}

func (a Admin) Login(login forms.AdminLogin) (token string, err error) {
	if err = DB.Where("name = ? ", login.UserName).First(&a).Error; gorm.IsRecordNotFoundError(err) {
		return "", err
	}
	if ok, _ := utils.ComparePasswords(a.Password, []byte(login.Password)); ok {
		expiresAt := time.Now().Add(time.Hour * time.Duration(72)).Unix()
		claims := tokenClaims{
			AdminID: strconv.Itoa(int(a.ID)),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expiresAt,
				Issuer:    "bandzest-auth",
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		c := config.GetConfig()
		tokenSecret := c.Get("jwt.token")
		if tokenString, err := token.SignedString([]byte(tokenSecret.(string))); err != nil {
			return "", err
		} else {
			return tokenString, nil
		}
	} else {
		return "", errors.New("invalid password")
	}
}

func (a Admin) Create(signup forms.AdminLogin) {
	password, _ := utils.HashAndSalt([]byte(signup.Password))
	admin := &Admin{
		Name:     signup.UserName,
		Password: password,
	}
	DB.Create(admin)
}
