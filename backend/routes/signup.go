package routes

import (
	"github.com/MishraShardendu22/controller"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SignupRoutes(app *fiber.App, collections *mongo.Collection) {
	app.Post("/signup", func(c *fiber.Ctx) error {
		return controllers.SignupHandler(c, collections)
	})
}
