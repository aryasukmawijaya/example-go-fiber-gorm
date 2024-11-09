package main

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/database/migration"
	"go-fiber-gorm/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DBInit()
	migration.RunMigration()

	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":3000")
}
