package routes

import (
	"fmt"
	"time"

	"github.com/MishraShardendu22/schema"
	"github.com/MishraShardendu22/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/exp/rand"
)

var otp int
var username_temp string

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
		if collections.FindOne(c.Context(), bson.M{"$or": []bson.M{
			{"email": UserSignUp.Email},
			{"username": UserSignUp.Username},
		}}).Decode(&UserSignUp) == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "User Already Exists",
			})
		}
		fmt.Println("Debug - 1")

		// Sending OTP
		rand.Seed(uint64(time.Now().UnixNano())) // Seed the random generator
		otp = rand.Intn(900000) + 100000
		fmt.Println("OTP:", otp)
		utils.SendEmailFast(UserSignUp.Email, otp)
		username_temp = UserSignUp.Username

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

func CheckOTP(app *fiber.App, collections *mongo.Collection) {
	app.Post("/checkotp", func(c *fiber.Ctx) error {
		fmt.Println("This is The Check OTP Route")

		var otpCheck struct {
			Val int `json:"val"` // Corrected the field name and added JSON tag
		}
		if err := c.BodyParser(&otpCheck); err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid JSON",
			})
		}

		// Assuming otp is a constant or a value retrieved from a secure source
		if otpCheck.Val != otp {
			collections.DeleteOne(c.Context(), bson.M{"username": username_temp})
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid OTP Authentication Failed",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   false,
			"message": "OTP Authentication Successful",
		})
	})
}
