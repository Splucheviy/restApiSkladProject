package main

import (
	"restApiSkladProject/cmd/apiserver"
	_ "restApiSkladProject/docs"
	"restApiSkladProject/pkg/db"
)

// @title			Sklad API Documentation
// @version		1.0
// @description	Sklad API Documentation
// @termsOfService	http://swagger.io/terms/
func main() {
	DB := db.Connect()
	db.CreateTable(DB)
	apiserver.HandleRequests(DB)
	db.CloseConnection(DB)
}
