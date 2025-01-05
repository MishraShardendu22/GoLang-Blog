package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type LoginDetails struct {
	Data     string `json:"data"`
	Password string `json:"pass"`
}

var jwtSecret = []byte("your-secret-key") // Replace with a secure secret

func GenerateToken(data string) (string, error) {
	claims := jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func VerifyToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}

func Login(app *fiber.App, collections *mongo.Collection) {
	app.Post("/login", func(c *fiber.Ctx) error {
		var userLogin LoginDetails
		if err := c.BodyParser(&userLogin); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid data format",
			})
		}

		var result bson.M
		err := collections.FindOne(c.Context(), bson.M{
			"$or": []bson.M{
				{"email": userLogin.Data},
				{"username": userLogin.Data},
			},
		}).Decode(&result)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid credentials",
			})
		}

		err = bcrypt.CompareHashAndPassword([]byte(result["password"].(string)), []byte(userLogin.Password))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid credentials",
			})
		}

		token, err := GenerateToken(userLogin.Data)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Could not generate token",
			})
		}

		fmt.Println(token)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   false,
			"message": "Login successful",
			"token":   token,
		})
	})
}

func AuthMiddleware(c *fiber.Ctx) error {
	tokenStr := c.Get("Authorization")
	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "No token provided",
		})
	}

	_, err := VerifyToken(tokenStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid token",
		})
	}

	return c.Next()
}
