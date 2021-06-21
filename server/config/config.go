package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"server/migration"
)

var (
	err    = godotenv.Load()
	mysqlUser = os.Getenv("MYSQL_USER")
	mysqlPass = os.Getenv("MYSQL_PASSWORD")
	mysqlHost = os.Getenv("MYSQL_HOST")
	mysqlName = os.Getenv("MYSQL_DATABASE")
)

func ConnectionToDatabase() *gorm.DB  {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPass, mysqlHost, mysqlName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&migration.User{}, &migration.List{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
