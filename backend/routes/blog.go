package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/MishraShardendu22/controllers"
)

func BlogRoutes(app *fiber.App, collections *mongo.Collection) {
	app.Post("/makeBlog", func(c *fiber.Ctx) error {
		return controllers.MakeBlog(c, collections)
	})
	app.Post("/deleteBlog", func(c *fiber.Ctx) error {
		return controllers.DeleteBlog(c, collections)
	})
	app.Put("/editBlog", func(c *fiber.Ctx) error {
		return controllers.EditBlog(c, collections)
	})
	app.Get("/getBlog", func(c *fiber.Ctx) error {
		return controllers.GetBlog(c, collections)
	})

}
