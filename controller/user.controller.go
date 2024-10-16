package controller

import (
	"learn-go/database"
	"learn-go/model/entity"
	"learn-go/model/request"
	"learn-go/model/response"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserControllerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)
}

func UserControllerGetByID(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, userID).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	userResponse := response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    userResponse,
	})
}

func UserControllerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}

func UserControllerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User
	userID := ctx.Params("id")

	err := database.DB.First(&user, userID).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "failed to update the data",
		})
	}

	//UPDATE DATA
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone
	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserControllerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserEmailRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User
	userID := ctx.Params("id")

	err := database.DB.First(&user, userID).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "failed to update the data",
		})
	}

	//UPDATE DATA
	user.Email = userRequest.Email

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserControllerDelete(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
	var user entity.User

	userRequest := new(request.UserEmailRequest)

	err := database.DB.Debug().First(&user, "id=?", userID).Error
	// CHECK AVAILABLE USER
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	//DELETE DATA

	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "user deleted",
	})
}
