package database

import (
	"log"
	"my-blog-site/server/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnctDB() {
	dsn := "root:Gauravjisri@21@tcp(127.0.0.1:3306)/fiber_blog?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Databse connection failed!")
	}

	log.Println("Connection Successful!")

	db.AutoMigrate(new(model.Blog))
	DBConn = db
}
