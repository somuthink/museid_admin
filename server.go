package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/somuthink/museid_admin/handlers"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/stats", func(c *fiber.Ctx) error {
		return c.Render("stats", fiber.Map{})
	})

	app.Get("/gtrooms", handlers.GtRoomsHX)

	app.Post("/dlRoom", handlers.DlRoomHX)

	app.Post("/upRoom", handlers.UpRoomHX)

	log.Fatal(app.Listen(":3000"))
}
