package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Blog(app *fiber.App, collections *mongo.Collection) {
	fmt.Println("This is The Blog Route")
}