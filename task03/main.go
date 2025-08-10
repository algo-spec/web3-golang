package main

import (
	// "task3/sql_base"

	"task3/sql_gorm"
	// "task3/sql_sqlx"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/jmoiron/sqlx"
	// _ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:12345@tcp(127.0.0.1:3307)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	// sql_base.Run(db)

	// db, err := sqlx.Connect("mysql", "root:12345@tcp(127.0.0.1:3307)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	// if err != nil {
	// 	panic(err)
	// }

	// sql_sqlx.Run(db)

	sql_gorm.Run(db)
}
