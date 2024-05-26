package main

import (
	"dating-apps/api/cmd/routes"
	"dating-apps/api/pkg/configs"
	"dating-apps/api/pkg/databases"
	"dating-apps/api/pkg/injectors"
)

func main() {
	db := databases.ConnectDatabase()
	validate := configs.InitValidate()
	service := injectors.InitService(db, validate)

	routes.Init(service)
}
