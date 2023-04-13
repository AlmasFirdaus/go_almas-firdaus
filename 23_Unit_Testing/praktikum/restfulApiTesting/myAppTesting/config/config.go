package config

import (
	"fmt"
	"myAppTesting/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() {

	config := models.Config{
		DB_Username: "root",
		DB_Password: "1234",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitialMigration()
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}

func InitDBTest() {

	configTest := models.ConfigTest{
		DB_Username_Test: "root",
		DB_Password_Test: "1234",
		DB_Port_Test:     "3306",
		DB_Host_Test:     "localhost",
		DB_Name_Test:     "crud_go_test",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configTest.DB_Username_Test, configTest.DB_Password_Test, configTest.DB_Host_Test, configTest.DB_Port_Test, configTest.DB_Name_Test)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitialMigrationTest()
}

func InitialMigrationTest() {
	DB.Migrator().DropTable(&models.User{})
	DB.AutoMigrate(&models.User{})
}
