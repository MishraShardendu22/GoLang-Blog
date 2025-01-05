package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeComment(c *fiber.Ctx, collections *mongo.Collection) error {
	// Dummy response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "MakeComment dummy function"})
}

func EditComment(c *fiber.Ctx, collections *mongo.Collection) error {
	// Dummy response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "DeleteComment dummy function"})
}

func DeleteComment(c *fiber.Ctx, collections *mongo.Collection) error {
	// Dummy response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "DeleteComment dummy function"})
}

func GetComment(c *fiber.Ctx, collections *mongo.Collection) error {
	// Dummy response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "GetComment dummy function"})
}
