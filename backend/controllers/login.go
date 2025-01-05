package controllers

import (
	"fmt"

	"github.com/MishraShardendu22/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type LoginDetails struct {
	Data     string `json:"data"`
	Password string `json:"pass"`
}

func LoginHandler(c *fiber.Ctx, collections *mongo.Collection) error {
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

	token, err := utils.GenerateToken(userLogin.Data)
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
}
