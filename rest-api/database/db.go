package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	DB_USER     = "root"
	DB_PASSWORD = ""
	DB_NAME     = "order_assignment"
	PARSE_TIME  = "true"
)

var (
	db  *sql.DB
	err error
)

func StartDB() {

	mysqlInfo := fmt.Sprintf("%s:%s@/%s?parseTime=%s", DB_USER, DB_PASSWORD, DB_NAME, PARSE_TIME)
	db, err = sql.Open("mysql", mysqlInfo)

	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	fmt.Println("Successfully connected to database")
}

func GetDBInstance() *sql.DB {
	return db
}
