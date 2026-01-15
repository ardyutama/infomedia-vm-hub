package controllers

import (
	"infomedia-vm-hub/database"
	"infomedia-vm-hub/handlers"
	"infomedia-vm-hub/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func AddContracts(c *fiber.Ctx) error {
	u := new(models.Contract)
	handlers.BodyParser(c, &u)
	handlers.AddEntity(c, &u)
	return c.Status(fiber.StatusCreated).JSON(u)
}

func GetContracts(c *fiber.Ctx) error {
	u := []models.Contract{}
	database.DB.Preload(clause.Associations).Find(&u)
	return c.JSON(u)
}
