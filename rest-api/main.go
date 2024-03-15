package main

import (
	"rest-api/controller"
	"rest-api/database"
	"rest-api/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var PORT = ":9090"

	// start db
	database.StartDB()

	db := database.GetDBInstance() // instance db

	controller.SetDBInstance(db) // set db for controller

	// call router
	router.StartOrderServer().Run(PORT)
}
