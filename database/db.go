package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	const MYSQL = "root:@tcp(127.0.0.1:3306)/go_fiber_gorm?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannote connect to database!")
	}

	fmt.Println("Successfully connected to database")
}
