package controllers

import (
	"infomedia-vm-hub/database"
	"infomedia-vm-hub/handlers"
	"infomedia-vm-hub/models"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllRoles(c *fiber.Ctx) error {
	roles := []models.Role{}
	database.DB.Find(&roles)
	return c.JSON(roles)
}

func AddRoles(c *fiber.Ctx) error {
	u := new(models.Role)
	handlers.BodyParser(c, u)
	if err := handlers.AddUniqueEntity(c, u); err != nil {
		// Handle the error response
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(&models.RoleResponse{
		Name: u.Name,
	})
}

func DeleteRoles(c *fiber.Ctx) error {
	role := new(models.Role)

	id := c.Params("id")

	err := database.DB.First(&role, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle the case where the record doesn't exist
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
		}
		// Handle other errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	if err := database.DB.Delete(&role).Error; err != nil {
		// Handle the error during deletion
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete record"})
	}

	database.DB.Where("id = ?", id).Delete(&role)

	return c.Status(200).JSON(fiber.Map{"message": "Successfully deleted"})
}

func EditRoles(c *fiber.Ctx) error {
	role := new(models.Role)

	handlers.BodyParser(c, &role)

	id := c.Params("id")

	database.DB.Model(&models.Role{}).Where("id = ?", id).Update("name", role.Name)

	return c.Status(400).JSON(fiber.Map{"message": "Successfully updated"})
}
