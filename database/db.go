package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	var err error

	viper.SetConfigFile(".env.toml")

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	db_host := viper.GetString("database.DB_HOST")
	db_port := viper.GetInt("database.DB_PORT")
	db_name := viper.GetString("database.DB_DATABASE")
	db_user := viper.GetString("database.DB_USERNAME")
	db_pass := viper.GetString("database.DB_PASSWORD")

	MYSQL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_host, db_port, db_name)

	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannote connect to database!")
	}

	fmt.Println("Successfully connected to database")
}
