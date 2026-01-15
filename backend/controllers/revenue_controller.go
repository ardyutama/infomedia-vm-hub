package controllers

import (
	"infomedia-vm-hub/database"
	"infomedia-vm-hub/models"

	"github.com/gofiber/fiber/v2"
)

func GetProjectRevenue(c *fiber.Ctx) error {
	u := []models.Project{}
	database.DB.Find(&u)
	responses := make([]models.ProjectDashboardResponse, len(u))

	// Map each project to a ProjectDashboardResponse object
	for i, project := range u {
		response := models.ProjectDashboardResponse{
			ID:          project.ID,
			ServiceName: project.ServiceName,
			ProjectName: project.ProjectName,
			Revenue:     project.Revenue,
			TotalVM:     project.TotalVM,
		}
		responses[i] = response
	}

	// Return the responses
	return c.JSON(responses)
}

func GetTotalRevenue(c *fiber.Ctx) error {
	var results struct {
		Total_vm int     `json:"total_vm"`
		Revenue  float64 `json:"revenue"`
	}

	database.DB.Table("projects").
		Select("COUNT(vm_specifications.id) as total_vm, SUM(CASE WHEN vm_specifications.vm_status_id = 1 THEN prices.active_running_price WHEN vm_specifications.vm_status_id = 2 THEN prices.active_stopped_price ELSE 0 END) as revenue").
		Joins("LEFT JOIN vm_specifications ON projects.id = vm_specifications.project_id").
		Joins("LEFT JOIN prices ON vm_specifications.id = prices.vm_specification_id").
		First(&results)

	// Print the results
	return c.JSON(&results)
}

func GetActiveRunningRevenue(c *fiber.Ctx) error {
	var results struct {
		Total_vm int     `json:"total_vm"`
		Revenue  float64 `json:"revenue"`
	}

	database.DB.Table("projects").
		Select("COUNT(vm_specifications.id) as total_vm, SUM(CASE WHEN vm_specifications.vm_status_id = 1 THEN prices.active_running_price ELSE 0 END) as revenue").
		Joins("LEFT JOIN vm_specifications ON projects.id = vm_specifications.project_id").
		Joins("LEFT JOIN prices ON vm_specifications.id = prices.vm_specification_id").
		Where("vm_specifications.vm_status_id = ?", 1).
		Scan(&results)

	// Print the results
	return c.JSON(&results)
}

func GetActiveStoppedRevenue(c *fiber.Ctx) error {
	var results struct {
		Total_vm int     `json:"total_vm"`
		Revenue  float64 `json:"revenue"`
	}

	database.DB.Table("projects").
		Select("COUNT(vm_specifications.id) as total_vm, SUM(CASE WHEN vm_specifications.vm_status_id = 2 THEN prices.active_stopped_price ELSE 0 END) as revenue").
		Joins("LEFT JOIN vm_specifications ON projects.id = vm_specifications.project_id").
		Joins("LEFT JOIN prices ON vm_specifications.id = prices.vm_specification_id").
		Where("vm_specifications.vm_status_id = ?", 2).
		Scan(&results)

	// Print the results
	return c.JSON(&results)
}

func GetSegmentRevenue(c *fiber.Ctx) error {
	type Result struct {
		Segment  string  `json:"segment"`
		TotalVMs int     `json:"total_vms"`
		Revenue  float64 `json:"revenue"`
	}

	var results []Result

	// Query to group by segment and calculate the total VMs and revenue for each segment
	database.DB.Table("projects").
		Select("segment, COUNT(vm_specifications.id) as TotalVMs, SUM(CASE WHEN vm_specifications.vm_status_id = 1 THEN prices.active_running_price WHEN vm_specifications.vm_status_id = 2 THEN prices.active_stopped_price ELSE 0 END) as revenue").
		Joins("LEFT JOIN vm_specifications ON projects.id = vm_specifications.project_id").
		Joins("LEFT JOIN prices ON vm_specifications.id = prices.vm_specification_id").
		Group("segment").
		Scan(&results)

	// Print the results
	return c.JSON(&results)
}
func GetHighestResourceUsage(c *fiber.Ctx) error {
	var result struct {
		Cpu           int    `json:"highest_cpu_count"`
		TotalCPU      int    `json:"total_cpu_vm"`
		Memory        int    `json:"highest_memory_count"`
		TotalMemory   int    `json:"total_memory_vm"`
		Hdd           int    `json:"highest_hdd_count"`
		TotalHdd      int    `json:"total_hdd_vm"`
		DiskType      string `json:"highest_disk_type_count"`
		TotalDiskType int    `json:"total_disk_type_vm"`
	}

	// Query to get the highest CPU count and the total number opu using that count
	database.DB.Table("vm_specifications").
		Select("cpu, COUNT(id) as TotalCPU").
		Group("cpu").
		Order("TotalCPU DESC").
		Limit(1).
		Scan(&result)

	// Query to get the highest memory size and the total number opu using that size
	database.DB.Table("vm_specifications").
		Select("Memory, COUNT(id) as TotalMemory").
		Group("Memory").
		Order("TotalMemory DESC").
		Limit(1).
		Scan(&result)

	// Query to get the highest HDD size and the total number of VMs using that size
	database.DB.Table("vm_specifications").
		Select("Hdd, COUNT(id) as TotalHdd").
		Group("Hdd").
		Order("TotalHdd DESC").
		Limit(1).
		Scan(&result)

	// Query to get the highest disk type and the total number of VMs using that disk type
	database.DB.Table("vm_specifications").
		Select("disk_types.name as DiskType, COUNT(vm_specifications.id) as TotalDiskType").
		Joins("LEFT JOIN disk_types ON vm_specifications.disk_type_id = disk_types.id").
		Group("DiskType").
		Order("TotalDiskType DESC").
		Limit(1).
		Scan(&result)

	// Print the results
	return c.JSON(&result)
}

func GetLocationCount(c *fiber.Ctx) error {
	type Result struct {
		Location string `json:"location"`
		TotalVMs int    `json:"total_vms"`
	}

	var results []Result

	// Query to group by segment and calculate the total VMs and revenue for each segment
	database.DB.Table("projects").
		Select("site_locations.name as Location, COUNT(vm_specifications.id) as TotalVMs").
		Joins("LEFT JOIN vm_specifications ON projects.id = vm_specifications.project_id").
		Joins("LEFT JOIN site_locations ON projects.site_location_id = site_locations.id").
		Group("Location").
		Scan(&results)

	// Print the results
	return c.JSON(&results)
}
