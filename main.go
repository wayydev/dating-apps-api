package main

import (
	"dating-apps/api/cmd/routes"
	"dating-apps/api/pkg/databases"
	"dating-apps/api/pkg/injectors"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := databases.ConnectDatabase()
	validate := validator.New()
	service := injectors.InitService(db, validate)
	routes.Init(service)
}
