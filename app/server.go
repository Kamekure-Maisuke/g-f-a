package main

import "github.com/gofiber/fiber"

func main() {
	// setting
	app := fiber.New(&fiber.Settings{
		TemplateEngine: "amber",
		TemplateFolder: "views",
		TemplateExtension: ".amber",
	})

	// Group
	apiGroup := app.Group("/api/v1")

	// index
  app.Get("/", func(c *fiber.Ctx) {
		data := fiber.Map{
			"name": "satou",
			"age":  34,
		}
		c.Render("index",data)
  })
	// param
	apiGroup.Get("/:id",func(c *fiber.Ctx) {
		c.Send("id = " + c.Params("id"))
	})

  app.Listen(3000)
}