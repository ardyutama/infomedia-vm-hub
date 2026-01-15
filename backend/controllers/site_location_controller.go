package controllers

import (
	"infomedia-vm-hub/database"
	"infomedia-vm-hub/handlers"
	"infomedia-vm-hub/models"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetSiteLocations(c *fiber.Ctx) error {
	u := []models.SiteLocation{}
	database.DB.Find(&u)
	return c.JSON(u)
}
func AddSiteLocations(c *fiber.Ctx) error {
	u := new(models.SiteLocation)
	handlers.BodyParser(c, &u)
	if err := handlers.AddUniqueEntity(c, &u); err != nil {
		// Handle the error response
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(&models.SiteLocationResponse{
		Name: u.Name,
	})
}

func DeleteSiteLocations(c *fiber.Ctx) error {
	u := new(models.SiteLocation)

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

func EditSiteLocations(c *fiber.Ctx) error {
	u := new(models.SiteLocation)

	if err := c.BodyParser(u); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	id := c.Params("id")

	err := database.DB.Model(&models.SiteLocation{}).Where("id = ?", id).Update("name", u.Name).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle the case where the record doesn't exist
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
		}
		// Handle other errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.Status(400).JSON(fiber.Map{"message": "Successfully Updated"})
}
