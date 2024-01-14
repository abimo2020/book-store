package config

import (
	"fmt"

	"github.com/abimo2020/book-store/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type DBConfig struct {
	Username string
	Password string
	Port     string
	Host     string
	Name     string
}

func InitDB() {
	utils.LoadEnv()

	dbConfig := DBConfig{
		Username: utils.GetEnv("DB_USERNAME", "root"),
		Password: utils.GetEnv("DB_PASSWORD", "popo1212"),
		Port:     utils.GetEnv("DB_PORT", "3306"),
		Host:     utils.GetEnv("DB_HOST", "localhost"),
		Name:     utils.GetEnv("DB_Name", "go-book-store"),
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	var err error
	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitMigration() {
	db.AutoMigrate()
}
