package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	)

var conn = &sql.DB{}

func ConnDb()(err error) {
	conn, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/way?tls=skip-verify&autocommit=true")
	if err != nil {
		return err
	}
	err = conn.Ping()
	if err != nil {
		return err
	}
	conn.SetMaxOpenConns(11)
	return nil
}

func InsertRows(data string)(err error) {
	fmt.Println("INSERT INTO qq_mobile_1 (id,qq ,mobile)VALUES" + data)
	_, err = conn.Exec("INSERT INTO qq_mobile_1 (id,qq ,mobile)VALUES ('%s', '%s', '%s')",  data)
	if err != nil {
		return err
	}
	return
}
