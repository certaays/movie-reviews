// package main

// import (
// 	"learn-go/database"
// 	"learn-go/migration"
// 	"learn-go/route"

// 	"github.com/gofiber/fiber/v2"
// )

// func Main() {
// 	//INITIAL DATABASE
// 	database.DatabaseInit()

// 	//RUN MIGRATION
// 	migration.RunMigration()

// 	app := fiber.New()

// 	//INITIAL ROUTE
// 	route.RouteInit(app)

// 	app.Listen(":8080")
// }

// func main() {
// 	Main()
// }

package handler

import (
	"learn-go/database"
	"learn-go/migration"
	"learn-go/route"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Main() *fiber.App {
	// INITIAL DATABASE
	database.DatabaseInit()

	// RUN MIGRATION
	migration.RunMigration()

	app := fiber.New()

	// INITIAL ROUTE
	route.RouteInit(app)

	return app
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app := Main()
	app.Handler()(w, r)
}
