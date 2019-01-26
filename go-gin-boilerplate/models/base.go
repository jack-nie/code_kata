package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/config"
	"github.com/sirupsen/logrus"

	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var DB *gorm.DB

type BaseModel struct {
	Id        int32
	DeletedAt *time.Time
}

func (m BaseModel) NewRecord() bool {
	return m.Id <= 0
}

func (m BaseModel) Destroy() error {
	err := db.Delete(&m).Error
	return err
}

func (m BaseModel) IsDeleted() bool {
	return m.DeletedAt != nil
}

func InitDatabase() {
	c := config.GetConfig()
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s", c.Get("mysql.username"), c.Get("mysql.password"), c.Get("mysql.dbname")))
	DB = db
	if err != nil {
		logrus.Fatalf("connect to mysql failed: %s", err)
	}

	db.AutoMigrate(&User{})
}
