package main

import (
	"fmt"
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	userName := "root"
	password := "111"
	host := "127.0.0.1"
	port := 3306
	DbName := "ginchat"
	timeout := "10s"

	// newLogger := logger.Default.LogMode(logger.Info)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", userName, password, host, port, DbName, timeout)
    
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// SkipDefaultTransaction: false,
		// NamingStrategy: schema.NamingStrategy{
		// 	TablePrefix:   "", 
		// 	SingularTable: false,
		// 	NoLowerCase:   false,
		// },
		// Logger: newLogger,
	})
	if err != nil {
		panic("connect db error, error = " + err.Error())
	}
	db.AutoMigrate(&models.Message{}) 
    db.AutoMigrate(&models.Contact{}) 
    db.AutoMigrate(&models.GroupBasic{}) 
    // db.AutoMigrate(&models.Contact{}) 

	// // Create
	// user := &models.UserBasic{}
	// user.Name = "flh"
	// db.Create(user)

	// // // Update
	// user.Email = "flh@163.com"
	// db.Save(user)

	// // Read
	// db.First(user)
	// fmt.Println(user)

	// // Delete
	// db.Delete(user)
}
