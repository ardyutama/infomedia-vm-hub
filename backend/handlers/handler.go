package handlers

import (
	"infomedia-vm-hub/database"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func BodyParser(c *fiber.Ctx, entity interface{}) error {
	if err := c.BodyParser(entity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	return nil
}
func AddUniqueEntity(c *fiber.Ctx, entity interface{}) error {
	if err := database.DB.Create(entity).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return fiber.NewError(fiber.StatusConflict, "{\"error\": \"Entity with the same name already exists\"}")
		}
		// Handle other errors
		return fiber.NewError(fiber.StatusInternalServerError, "{\"error\": \"An unexpected error occurred\"}")
	}
	// Return entity in JSON format
	return nil

}
func AddEntity(c *fiber.Ctx, entity interface{}) error {
	if err := database.DB.Create(entity).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "An unexpected error occurred",
		})
	}

	// Return entity in JSON format
	return c.Status(fiber.StatusCreated).JSON(entity)
}
