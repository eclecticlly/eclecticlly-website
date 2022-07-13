package router

import (
	"eclecticlly/web-stack/app/controllers"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

type HttpRouter struct {
}

func (h HttpRouter) InstallRouter(app *fiber.App) {
	group := app.Group("", cors.New(), csrf.New(csrf.Config{
		// only to control the switch whether csrf is activated or not
		Next: func(c *fiber.Ctx) bool {
			return true
		},
		KeyLookup:      "form:_csrf",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
		ContextKey:     "token",
	}))

	group.Get("/", controllers.RenderHello)
	group.Post("/clicked", controllers.RenderClicked)
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
