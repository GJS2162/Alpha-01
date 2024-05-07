package controller

import (
	"github.com/gofiber/fiber/v2"
)

// Blog list
func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog List",
	}

	c.Status(200)
	return c.JSON(context)
}

// Add a blog in databse
func BlogCreate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Delete blog for given ID",
	}

	c.Status(200)
	return c.JSON(context)
}

// Update a blog
func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a blog",
	}

	c.Status(201)
	return c.JSON(context)
}

// Delete a blog
func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update a blog",
	}

	c.Status(200)
	return c.JSON(context)
}
