package controllers

import "github.com/gofiber/fiber/v2"

func RenderHello(c *fiber.Ctx) error {

	return c.Render("index.jet", fiber.Map{
		"FiberTitle": "Hello From Fiber Jet Engine!",
	}, "layouts/main.jet")

}

func RenderClicked(c *fiber.Ctx) error {
	return c.Render("htmx.jet", nil)
}
