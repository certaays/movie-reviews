package route

import (
	"learn-go/config"
	"learn-go/controller"
	"learn-go/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("public", config.ProjectRootPath+"/public/assets")
	r.Post("/login", controller.Login)
	r.Get("/user", middleware.Auth, controller.UserControllerGetAll)
	r.Get("/user/:id", controller.UserControllerGetByID)
	r.Post("/user", controller.UserControllerCreate)
	r.Put("/user/:id", controller.UserControllerUpdate)
	r.Put("/user/:id/update-email", controller.UserControllerUpdateEmail)
	r.Delete("/user/:id", controller.UserControllerDelete)
}
