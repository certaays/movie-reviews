package main

import (
	"learn-go/database"
	"learn-go/migration"
	"learn-go/route"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vercel/vercel-go"
)

func handler(w http.ResponseWriter, r *http.Request) {
	database.DatabaseInit()
	migration.RunMigration()
	app := fiber.New()
	route.RouteInit(app)
	app.Handler(w, r)
}

func main() {
	vercel.Start(handler)
}
