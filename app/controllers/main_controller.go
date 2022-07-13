package controllers

import "github.com/gofiber/fiber/v2"

// RenderHello is the basic home route
func RenderHello(c *fiber.Ctx) error {

	return c.Render("views/index.jet", fiber.Map{
		"FiberTitle": "Hello From Fiber Jet Engine!",
	}, "views/layouts/main.jet")

}

// Renderclicked is the test to see if htmx works
func RenderClicked(c *fiber.Ctx) error {
	return c.Render("views/htmx.jet", nil)
}
