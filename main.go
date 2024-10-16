package main

import (
	"learn-go/database"
	"learn-go/migration"
	"learn-go/route"

	"github.com/gofiber/fiber/v2"
)

func Main() {
	//INITIAL DATABASE
	database.DatabaseInit()

	//RUN MIGRATION
	migration.RunMigration()

	app := fiber.New()

	//INITIAL ROUTE
	route.RouteInit(app)

	app.Listen(":8080")
}
