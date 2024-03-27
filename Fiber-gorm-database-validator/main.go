package main

import (
	"github.com/gofiber/fiber/v2"
	"gofiberind/database"
	"gofiberind/models/entity"
	"gofiberind/routers"
)

func main() {
	database.ConnectDB()
	migration.RunMigrate()
	app := fiber.New()
	routers.RouterApp(app)
	app.Listen(":8080")
}
