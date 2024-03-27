package routers

import (
	"github.com/gofiber/fiber/v2"
	"gofiberind/controllers"
)

func RouterApp(c *fiber.App) {
	c.Get("/api/showall", controllers.UserControllerShow)
	c.Get("/api/userGetbyId/:id", controllers.UserControllerGetById)
	c.Post("/api/create", controllers.UserControllerCreate)
}
