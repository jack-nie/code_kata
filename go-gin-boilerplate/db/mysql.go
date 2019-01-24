package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/config"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

func Init() {
	var err error
	c := config.GetConfig()
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", c.Get("mysql.username"), c.Get("mysql.password"), c.Get("mysql.dbname")))
	if err != nil {
		logrus.Fatalf("connect to mysql failed: %s", err)
	}
}

func GetDB() *sql.DB {
	return db
}
