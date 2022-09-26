package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("connection failed.")
	}

	return db
}

func Close(gd *gorm.DB) {
	sqlDB, _ := gd.DB()
	sqlDB.Close()
}
