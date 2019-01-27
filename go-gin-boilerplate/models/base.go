package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/config"
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var DB *gorm.DB

type BaseModel struct {
	gorm.Model
}

func (m BaseModel) NewRecord() bool {
	return m.ID <= 0
}

func (m BaseModel) Destroy() error {
	err := db.Delete(&m).Error
	return err
}

func (m BaseModel) IsDeleted() bool {
	return m.DeletedAt != nil
}

type GormLogger struct{}

func (*GormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		logrus.WithFields(logrus.Fields{"module": "gorm", "type": "sql"}).Print(v[3])
	}
	if v[0] == "log" {
		logrus.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}

func InitDatabase() {
	c := config.GetConfig()
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s", c.Get("mysql.username"), c.Get("mysql.password"), c.Get("mysql.dbname")))
	DB = db
	if err != nil {
		logrus.Fatalf("connect to mysql failed: %s", err)
	}
	db.SetLogger(&GormLogger{})
	db.LogMode(true)

	db.AutoMigrate(&User{}, &Author{}, &Category{}, &Comment{}, &Post{}, &Tag{})
	db.Model(&Post{}).AddIndex("index_on_category_id", "category_id")
	db.Model(&Post{}).AddIndex("index_on_author_id", "author_id")
	db.Model(&Comment{}).AddIndex("index_on_user_id", "user_id")
	db.Model(&Comment{}).AddIndex("index_on_post_id", "post_id")
}
