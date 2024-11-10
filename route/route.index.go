package route

import (
	"go-fiber-gorm/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/users", handler.UserHandlerGetAll)
	r.Post("/user", handler.UserHandlerCreate)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Put("/user/email/:id", handler.UserHandlerUpdateEmail)
	r.Put("/user/:id", handler.UserHandlerUpdate)
}
