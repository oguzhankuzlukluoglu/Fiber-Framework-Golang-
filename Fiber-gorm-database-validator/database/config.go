package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	const MYSQL = "root:burayasifregelecek@tcp(127.0.0.1:3306)/databaseismiburayagelecek?charset=utf8mb4&parseTime=True&loc=Local"
	DSN := MYSQL
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("can't connected to database")
	}
	fmt.Println("Connected to Database")
}
