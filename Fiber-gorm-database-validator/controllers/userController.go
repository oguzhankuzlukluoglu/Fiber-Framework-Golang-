package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gofiberind/database"
	"gofiberind/database/migration"
	"gofiberind/models/req"
	"log"
)

func UserControllerShow(c *fiber.Ctx) error {
	var users []entity.User
	err := database.DB.Find(&users).Error
	if err != nil {
		log.Println(err)
	}
	return c.JSON(users)
}
func UserControllerCreate(c *fiber.Ctx) error {
	user := new(req.UserReq)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	validation := validator.New()
	if err := validation.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed input new user",
			"error":   err.Error(),
		})
	}
	newUser := entity.User{
		Name:  user.Name,
		Email: user.Email,
	}
	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed create new user",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success to create new user",
		"data":    newUser,
	})
}
func UserControllerGetById(c *fiber.Ctx) error {
	var user []entity.User
	id := c.Params("id")
	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id boş olamaz",
		})
		return nil
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user verisi başarıyla alındı",
		"data":    user,
	})
}
