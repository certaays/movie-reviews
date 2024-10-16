package controller

import (
	"learn-go/database"
	"learn-go/model/entity"
	"learn-go/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}
	var user entity.User
	err := database.DB.First(&user, "email=?", loginRequest.Email).Error

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": "secret",
	})
}
