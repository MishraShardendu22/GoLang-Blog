package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeBlog(c *fiber.Ctx, collections *mongo.Collection) error {
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "MakeBlog dummy function"})
}

func DeleteBlog(c *fiber.Ctx, collections *mongo.Collection) error {
	// Dummy response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "DeleteBlog dummy function"})
}

func EditBlog(c *fiber.Ctx, collections *mongo.Collection) error {
	// Dummy response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "EditBlog dummy function"})
}

func GetBlog(c *fiber.Ctx, collections *mongo.Collection) error {
	// Dummy response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "GetBlog dummy function"})
}
