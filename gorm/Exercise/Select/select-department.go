package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type departement struct {
	id              uint
	nama_department string
	MemberNumber    sql.NullString
	ActivatedAt     sql.NullTime
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func main() {
	res := "host=localhost  user=postgres password=annisa073 dbname=db_annisa_nur port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(res), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

		//create
		db.Select("nama_department").Create(&departement{})

		if err := db.Error; err != nil {
			fmt.Println(err)
		}

		fmt.Println("insert Success")
	}
}
