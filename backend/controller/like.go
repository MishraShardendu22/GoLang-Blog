package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func LikePost(c *fiber.Ctx, collections *mongo.Collection) error {
	// Dummy response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "LikePost dummy function"})
}

func GetLikes(c *fiber.Ctx, collections *mongo.Collection) error {
	// Dummy response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "LikePost dummy function"})
}

func UnLikePost(c *fiber.Ctx, collections *mongo.Collection) error {
	// Dummy response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "UnLikePost dummy function"})
}

func LikedPost(c *fiber.Ctx, collections *mongo.Collection) error {
	// Dummy response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "LikedPost dummy function"})
}
