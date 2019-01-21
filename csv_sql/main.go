package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := sql.Open("mysql", "jack:jack@/test?charset=utf8")
	if err != nil {
		logrus.Fatalf("open mysql failed: %v", err)
	}

	stmt, err := db.Prepare("INSERT userinfo SET username=?, department=?, created=?")

	if err != nil {
		logrus.Fatalf("%v", err)
	}

	res, err := stmt.Exec("jack", "Sales", "2019-01-20")

	if err != nil {
		logrus.Fatalf("%v", err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		logrus.Fatalf("%v", err)
	}

	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	res, err = stmt.Exec("jackupdate", id)
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	rows, err := db.Query("select * from userinfo")
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		if err != nil {
			logrus.Fatalf("%v", err)
		}

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	stmt, err = db.Prepare("delete from userinfo where uid=?")
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	res, err = stmt.Exec(id)
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	affect, err = res.RowsAffected()
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	fmt.Println(affect)
	db.Close()
}
