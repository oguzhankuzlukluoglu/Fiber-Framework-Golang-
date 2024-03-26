package main

import (
	"firstAPI/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()
	//app.Use(basicauth.New(basicauth.Config{
	//	Users: map[string]string{
	//		"john":  "doe",
	//		"admin": "123456",
	//	},
	//}))
	app.Get("/customers", handlers.GetAllCustomer)
	app.Get("/customer/:username", handlers.GetByUsernameWithData)
	app.Post("/customer", handlers.CreateCustomer)
	app.Post("/customer/update", handlers.UpdateCustomer)
	app.Delete("/customer/:username", basicauth.New(basicauth.Config{Users: map[string]string{
		"admin": "964277",
	}}), handlers.DeleteCustomer)

	app.Listen(":8000")

}
