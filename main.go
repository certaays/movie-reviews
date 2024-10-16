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
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vercel/go-bridge/go/bridge"
)

var app *fiber.App

func init() {
	// INITIAL DATABASE
	database.DatabaseInit()

	// RUN MIGRATION
	migration.RunMigration()

	app = fiber.New()

	// INITIAL ROUTE
	route.RouteInit(app)
}

// Handler function to be called by Vercel
func Handler(ctx *bridge.RequestCtx) {
	if err := app.Handler()(ctx); err != nil {
		log.Printf("Error handling request: %v", err)
	}
}
