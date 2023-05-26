package main

import (
	// "crypto/sha256"
	"database/sql"
	"fmt"
	// "io/ioutil"
	// "log"

	_ "github.com/lib/pq"
)

type DBCredential struct {
	HostName     string
	DatabaseName string
	Username     string
	Password     string
	Port         string
}

var (
	// CAMP_ID = "BE1234567" // TODO: replace this

	credential = DBCredential{
		HostName:     "localhost",
		DatabaseName: "db_annisa_nur",
		Username:     "postgres",
		Password:     "annisa073",
		Port:         "5432",
		// TODO: answer here
	}
)

func Connect() (db *sql.DB, err error) {
	//setup connection to database postgres
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	credential.HostName, credential.Port, credential.Username, credential.Password, credential.DatabaseName)

	// TODO: answer here

	db, err = sql.Open("postgres", dns)
	if err != nil {
		return
	}

	return db, nil
}

func main() {
	db, err := Connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`INSERT INTO employess VALUES 
	(1, "Rizki", 25, "Jl. Kebon Jeruk, 2000000),
	)`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Table Employee Created!")
}