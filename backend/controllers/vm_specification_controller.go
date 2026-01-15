package controllers

import (
	"infomedia-vm-hub/database"
	"infomedia-vm-hub/handlers"
	"infomedia-vm-hub/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GetAllVMSpecifications(c *fiber.Ctx) error {
	// u := []models.VMSpecification{}
	var result []models.VMSpecificationResponse

	database.DB.Table("vm_specifications").
		Preload("Contract.ContractRequestType").
		Preload("Contract.RequestBasedType").
		Preload(clause.Associations).
		Select("vm_specifications.*, service_types.name AS service_type, vm_types.name AS vm_type, disk_types.name AS disk_type, operating_systems.name AS operating_system, vm_statuses.name AS vm_status, projects.project_name AS project_name, projects.service_name AS service_name, projects.segment AS segment, site_locations.name AS site_location, projects.gl_account, projects.cost_center").
		Joins("LEFT JOIN service_types ON service_types.id = vm_specifications.service_type_id").
		Joins("LEFT JOIN vm_types ON vm_types.id = vm_specifications.vm_type_id").
		Joins("LEFT JOIN disk_types ON disk_types.id = vm_specifications.disk_type_id").
		Joins("LEFT JOIN operating_systems ON operating_systems.id = vm_specifications.operating_system_id").
		Joins("LEFT JOIN vm_statuses ON vm_statuses.id = vm_specifications.vm_status_id").
		Joins("LEFT JOIN projects ON projects.id = vm_specifications.project_id").
		Joins("LEFT JOIN site_locations ON site_locations.id = projects.site_location_id").
		Find(&result)
	return c.JSON(&result)
}

func GetVMSpecificationsByID(c *fiber.Ctx) error {
	// u := []models.VMSpecification{}

	projectID := c.Params("id")

	var result models.VMSpecificationResponse

	database.DB.Table("vm_specifications").
		Preload("Contract.ContractRequestType").
		Preload("Contract.RequestBasedType").
		Preload(clause.Associations).
		Select("vm_specifications.*, service_types.name AS service_type, vm_types.name AS vm_type, disk_types.name AS disk_type, operating_systems.name AS operating_system, vm_statuses.name AS vm_status, projects.project_name AS project_name, projects.service_name AS service_name, projects.segment AS segment, site_locations.name AS site_location, projects.gl_account, projects.cost_center").
		Joins("LEFT JOIN service_types ON service_types.id = vm_specifications.service_type_id").
		Joins("LEFT JOIN vm_types ON vm_types.id = vm_specifications.vm_type_id").
		Joins("LEFT JOIN disk_types ON disk_types.id = vm_specifications.disk_type_id").
		Joins("LEFT JOIN operating_systems ON operating_systems.id = vm_specifications.operating_system_id").
		Joins("LEFT JOIN vm_statuses ON vm_statuses.id = vm_specifications.vm_status_id").
		Joins("LEFT JOIN projects ON projects.id = vm_specifications.project_id").
		Joins("LEFT JOIN site_locations ON site_locations.id = projects.site_location_id").
		Find(&result, projectID)

	return c.JSON(&result)
}

func GetVMSpecificationsByProjectID(c *fiber.Ctx) error {
	// u := []models.VMSpecification{}

	projectID := c.Params("id")

	var result []models.VMSpecificationResponse

	database.DB.Table("vm_specifications").
		Where("vm_specifications.project_id = ?", projectID).
		Preload("Contract.ContractRequestType").
		Preload("Contract.RequestBasedType").
		Preload(clause.Associations).
		Select("vm_specifications.*, service_types.name AS service_type, vm_types.name AS vm_type, disk_types.name AS disk_type, operating_systems.name AS operating_system, vm_statuses.name AS vm_status, projects.project_name AS project_name, projects.service_name AS service_name, projects.segment AS segment, site_locations.name AS site_location, projects.gl_account, projects.cost_center").
		Joins("LEFT JOIN service_types ON service_types.id = vm_specifications.service_type_id").
		Joins("LEFT JOIN vm_types ON vm_types.id = vm_specifications.vm_type_id").
		Joins("LEFT JOIN disk_types ON disk_types.id = vm_specifications.disk_type_id").
		Joins("LEFT JOIN operating_systems ON operating_systems.id = vm_specifications.operating_system_id").
		Joins("LEFT JOIN vm_statuses ON vm_statuses.id = vm_specifications.vm_status_id").
		Joins("LEFT JOIN projects ON projects.id = vm_specifications.project_id").
		Joins("LEFT JOIN site_locations ON site_locations.id = projects.site_location_id").
		Find(&result)

	return c.JSON(&result)
}

func AddVMSpecifications(c *fiber.Ctx) error {
	u := new(models.VMSpecification)
	handlers.BodyParser(c, &u)
	if err := database.DB.Create(&u).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "An unexpected error occurred",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(u)
}

// func AddVMSpecifications(c *fiber.Ctx) error {
// 	u := new(models.VMSpecification)
// 	if err := c.BodyParser(u); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	price := u.Price
// 	res := database.DB.Create(&price)

// 	if res.Error != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": res.Error.Error()})
// 	}

// 	u.PriceID = price.ID

// 	result := database.DB.Create(&u)

// 	if result.Error != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(u)

// }
