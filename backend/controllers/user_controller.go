package controllers

import (
	"infomedia-vm-hub/database"
	"infomedia-vm-hub/handlers"
	"infomedia-vm-hub/models"
	"infomedia-vm-hub/utils/password"
	"errors"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SignUp(c *fiber.Ctx) error {
	var u models.User

	// Parse body into struct
	handlers.BodyParser(c, &u)

	user := &models.User{
		Username: u.Username,
		Password: password.Generate(u.Password),
		RoleID:   u.RoleID,
	}

	// Return Employee in JSON format
	if err := database.DB.Create(user).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Entity with the same name already exists",
			})
		}
		// Handle other errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "An unexpected error occurred",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(models.UserResponse{Username: user.Username})
}

func Login(c *fiber.Ctx) error {
	var b models.User

	// Parse body into struct
	if err := c.BodyParser(&b); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(400).SendString("An error occurred while processing the request.")
	}

	u := &models.UserResponse{}

	err := database.DB.Model(&models.User{}).Where("username = ?", b.Username).First(&u).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Invalid Username/Password", err)
		return fiber.NewError(fiber.StatusUnauthorized, "Username Not Found")
	}

	if err := password.Verify(u.Password, b.Password); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Wrong Password")
	}

	return c.JSON(models.AuthResponse{Username: u.Username})
}
