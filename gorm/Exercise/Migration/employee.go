package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	
)

type Employee struct {
	id           	uint
	name         	string
	address        	string
	age          	uint8
	birthdate     	string
	level 			string 
	id_department	uint
	
  }

  func main() {
	res := "host=localhost  user=postgres password=annisa073 dbname=db_annisa_nur port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(res), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

		//migrate
		db.AutoMigrate(&Employee{})

		fmt.Println("Migration Success")
}