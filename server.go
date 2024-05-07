package main

import (
	"my-blog-site/server/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnctDB()

	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in sql connection!")
	}

	defer sqlDb.Close()

	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello World!!!")
	// })

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{"mesaage": "Welcome to the Automation World!"})
	// })

	// app.Get("/about", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{"message": "Welcome to about section!"})
	// })

	router.setupRoutes(app)
	app.Listen(":8000")
}
