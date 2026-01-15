// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
package main

import (
	"infomedia-vm-hub/database"
	"infomedia-vm-hub/models"
	"infomedia-vm-hub/route"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	database.Migrate(
		&models.Role{},
		&models.User{},
		&models.ServiceType{},
		&models.VMType{},
		&models.DiskType{},
		&models.Network{},
		&models.OperatingSystem{},
		&models.VMStatus{},
		&models.ContractRequestType{},
		&models.RequestBasedType{},
		&models.Contract{},
		&models.ContractHistory{},
		&models.Price{},
		&models.PriceHistory{},
		&models.VMSpecificationHistory{},
		&models.Contact{},
		&models.AdditionalFeature{},
		&models.ObjectStorage{},
		&models.SiteLocation{},
		&models.Project{},
		&models.VMSpecification{},
		&models.ResourcePerHour{},
	)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))

	route.ServiceRoutes(app)

	app.Listen(":8080")
}
