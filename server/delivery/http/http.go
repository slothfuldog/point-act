package http

import (
	"point/delivery/http/router"
	"point/infrastructure"

	"github.com/gofiber/fiber/v2"
)

func NewHttpDelivery(dir string) *fiber.App {
	app := fiber.New()
	var c *fiber.Ctx

	db := infrastructure.NewDatabaseConnect(dir)

	router.NewRouter(app, db, c)

	return app
}
