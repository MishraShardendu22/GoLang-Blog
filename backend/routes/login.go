package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type LoginDetails struct {
	Data     string `json:"data"`
	Password string `json:"pass"`
}

func Login(app *fiber.App, collections *mongo.Collection) {
	fmt.Println("This is The Login Route")
	app.Post("/login", func(c *fiber.Ctx) error {
		var userLogin LoginDetails

		if err := c.BodyParser(&userLogin); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "Data Not Entered",
			})
		}
		// Check 1
		// fmt.Println(userLogin)

		fmt.Println("Debug - 0")

		var result bson.M
		test := collections.FindOne(c.Context(), bson.M{
			"$and": []bson.M{
				{
					"$or": []bson.M{
						{"email": userLogin.Data},
						{"username": userLogin.Data},
					},
				},
			},
		}).Decode(&result)

		fmt.Println("Debug - 0.5", test)

		err := bcrypt.CompareHashAndPassword([]byte(result["password"].(string)), []byte(userLogin.Password))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid credentials",
			})
		}

		result["password"] = "hidden"
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   false,
			"message": "Logged In Successfully",
			"data":    result,
		})
	})
}
