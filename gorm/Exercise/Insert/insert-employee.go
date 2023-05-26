package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Employee struct {
	id           	uint
	name         	string
	address        	string
	age          	uint8
	birthdate     	string
	level 			string 
	id_department	uint
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
  }

  func main() {
	res := "host=localhost  user=postgres password=annisa073 dbname=db_annisa_nur port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(res), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

		//create
		db.Create(&Employee{id: 123, name: "annisa", address: "Jl.Kamboja", age: 21, birthdate: "2002-09-29", level: "Supervisor", id_department: 101, MemberNumber: sql.NullString{String: "123", Valid: true}, ActivatedAt: sql.NullTime{Time: time.Now(), Valid: true}})

		if err := db.Error; err != nil {
			fmt.Println(err)
		}

		fmt.Println("insert Success")
	}
}
