package main

import (
	"example/go-orm-api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func main() {
	dsn := "mitsu:secret@tcp(db:3306)/go_local?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Create
	db.Create(&model.User{
		Fname:"A1",
		Lname: "L1",
		Username: "A1@mecallapi.com",
		Avatar: "https://www.mecallapi.com/users/1.png",

	})
	// Create
	db.Create(&model.User{
		Fname:"A2",
		Lname: "L2",
		Username: "A2@mecallapi.com",
		Avatar: "https://www.mecallapi.com/users/2.png",

	})

	
}
