package controllers

import (
	"infomedia-vm-hub/database"
	"infomedia-vm-hub/handlers"
	"infomedia-vm-hub/models"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetVMTypes(c *fiber.Ctx) error {
	u := []models.VMType{}
	database.DB.Find(&u)
	return c.JSON(u)
}
func AddVMTypes(c *fiber.Ctx) error {
	u := new(models.VMType)
	handlers.BodyParser(c, &u)
	if err := handlers.AddUniqueEntity(c, &u); err != nil {
		// Handle the error response
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(&models.VMType{
		Name: u.Name,
	})
}

func DeleteVMTypes(c *fiber.Ctx) error {
	u := new(models.VMType)

	id := c.Params("id")

	err := database.DB.First(&u, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle the case where the record doesn't exist
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
		}
		// Handle other errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	if err := database.DB.Delete(&u).Error; err != nil {
		// Handle the error during deletion
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete record"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Successfully deleted"})
}

func EditVMTypes(c *fiber.Ctx) error {
	u := new(models.VMType)

	if err := c.BodyParser(u); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	id := c.Params("id")

	err := database.DB.Model(&models.VMType{}).Where("id = ?", id).Update("name", u.Name).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle the case where the record doesn't exist
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
		}
		// Handle other errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.Status(400).JSON(fiber.Map{"message": "Successfully updated"})
}
