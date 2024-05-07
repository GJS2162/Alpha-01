package gaurav

import (
	"my-blog-site/server/controller"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	//Get
	app.Get("/", controller.BlogList())

	//Post
	app.Post("/", controller.BlogCreate())

	//Put
	app.Put("/", controller.BlogUpdate())

	//Delete
	app.Delete("/", controller.BlogDelete())
}
