package controllers

import (
	"infomedia-vm-hub/database"
	"infomedia-vm-hub/handlers"
	"infomedia-vm-hub/models"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetNetworks(c *fiber.Ctx) error {
	u := []models.Network{}
	database.DB.Find(&u)
	return c.JSON(u)
}

func AddNetworks(c *fiber.Ctx) error {
	u := new(models.Network)
	handlers.BodyParser(c, &u)
	if err := handlers.AddUniqueEntity(c, &u); err != nil {
		// Handle the error response
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(&models.Network{
		IPPublic: u.IPPublic,
		IPLocal:  u.IPLocal,
		Port:     u.Port,
		VPCName:  u.VPCName,
	})
}

func DeleteNetworks(c *fiber.Ctx) error {
	u := new(models.Network)
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

func EditNetworks(c *fiber.Ctx) error {
	u := new(models.Network)

	if err := c.BodyParser(u); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	id := c.Params("id")

	exist := new(models.Network)
	res := database.DB.Model(&models.Network{}).Where("id = ?", id).First(exist)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			// Handle the case where the record doesn't exist
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
		}
		// Handle other errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	res = database.DB.Model(&models.Network{}).Where("id = ?", id).Updates(&u)

	if res.Error != nil {
		// Handle the error during update
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update record"})
	}

	return c.Status(400).JSON(fiber.Map{"message": "Successfully Updated"})
}
