package main

import "github.com/gofiber/fiber"

func main() {
	// setting
	app := fiber.New(&fiber.Settings{
		TemplateEngine: "amber",
		TemplateFolder: "views",
		TemplateExtension: ".amber",
	})
	// index
  app.Get("/", func(c *fiber.Ctx) {
		data := fiber.Map{
			"name": "john",
			"age":  34,
		}
		c.Render("index",data)
  })
	// param
	app.Get("/api/v1/:id",func(c *fiber.Ctx) {
		c.Send("id = " + c.Params("id"))
	})

  app.Listen(3000)
}