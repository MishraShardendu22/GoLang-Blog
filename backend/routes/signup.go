package routes

import (
	"fmt"

	"github.com/MishraShardendu22/schema"
	"github.com/MishraShardendu22/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Signup(app *fiber.App, collections *mongo.Collection) {
	app.Post("/signup", func(c *fiber.Ctx) error {
		fmt.Println("This is The Signup Route")

		var UserSignUp schema.User
		if err := c.BodyParser(&UserSignUp); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid JSON",
			})
		}
		fmt.Println("Debug - 0")

		// Set default values
		UserSignUp.SetDefaults()
		// fmt.Println(UserSignUp)

		fmt.Println("Debug - 0.5")
		// err := collections.FindOne(c.Context(), bson.M{"$or": []bson.M{
		// 	{"email": UserSignUp.Email},
		// 	{"username": UserSignUp.Username},
		// }}).Decode(&UserSignUp)

		// if err != nil {
		// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		// 		"error":   true,
		// 		"message": "User Already Exists",
		// 	})
		// }
		fmt.Println("Debug - 1")

		UserSignUp.Password = utils.HashPassWord(UserSignUp.Password)

		_, err := collections.InsertOne(c.Context(), UserSignUp)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "Error Registering The User Sorry!!",
			})
		}
		fmt.Println("Debug - 2")

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":   false,
			"message": "User Registered Successfully",
		})
	})
}
