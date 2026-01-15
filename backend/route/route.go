package route

import (
	"infomedia-vm-hub/controllers"

	"github.com/gofiber/fiber/v2"
)

func ServiceRoutes(app fiber.Router) {
	r := app.Group("/api/v1")

	r.Get("/roles", controllers.GetAllRoles)
	r.Post("/roles", controllers.AddRoles)
	r.Delete("/roles/:id", controllers.DeleteRoles)
	r.Put("/roles/:id", controllers.EditRoles)

	r.Post("/users/signup", controllers.SignUp)
	r.Post("/users/login", controllers.Login)

	r.Get("/service-types", controllers.GetServiceTypes)
	r.Post("/service-types", controllers.AddServiceTypes)
	r.Delete("/service-types/:id", controllers.DeleteServiceTypes)
	r.Put("/service-types/:id", controllers.EditServiceTypes)

	r.Get("/vm-types", controllers.GetVMTypes)
	r.Post("/vm-types", controllers.AddVMTypes)
	r.Delete("/vm-types/:id", controllers.DeleteVMTypes)
	r.Put("/vm-types/:id", controllers.EditVMTypes)

	r.Get("/site-locations", controllers.GetSiteLocations)
	r.Post("/site-locations", controllers.AddSiteLocations)
	r.Delete("/site-locations/:id", controllers.DeleteSiteLocations)
	r.Put("/site-locations/:id", controllers.EditSiteLocations)

	r.Get("/disk-types", controllers.GetDiskTypes)
	r.Post("/disk-types", controllers.AddDiskTypes)
	r.Delete("/disk-types/:id", controllers.DeleteDiskTypes)
	r.Put("/disk-types/:id", controllers.EditDiskTypes)

	r.Get("operating-systems", controllers.GetOperatingSystems)
	r.Post("operating-systems", controllers.AddOperatingSystems)
	r.Delete("operating-systems/:id", controllers.DeleteOperatingSystems)
	r.Put("operating-systems/:id", controllers.EditOperatingSystems)

	r.Get("vm-status", controllers.GetVMStatuses)
	r.Post("vm-status", controllers.AddVMStatuses)
	r.Delete("vm-status/:id", controllers.DeleteVMStatuses)
	r.Put("vm-status/:id", controllers.EditVMStatuses)

	r.Get("contract-request-types", controllers.GetContractRequestTypes)
	r.Post("contract-request-types", controllers.AddContractRequestTypes)
	r.Delete("contract-request-types/:id", controllers.DeleteContractRequestTypes)
	r.Put("contract-request-types/:id", controllers.EditContractRequestTypes)

	r.Get("request-based-types", controllers.GetRequestBasedTypes)
	r.Post("request-based-types", controllers.AddRequestBasedTypes)
	r.Delete("request-based-types/:id", controllers.DeleteRequestBasedTypes)
	r.Put("request-based-types/:id", controllers.EditRequestBasedTypes)

	r.Get("contacts", controllers.GetAllContacts)
	r.Post("contacts", controllers.AddContacts)
	r.Delete("contacts/:id", controllers.DeleteContacts)
	r.Put("contacts/:id", controllers.EditContacts)

	r.Get("networks", controllers.GetNetworks)
	r.Post("networks", controllers.AddNetworks)
	r.Delete("networks/:id", controllers.DeleteNetworks)
	r.Put("networks/:id", controllers.EditNetworks)

	r.Get("contracts", controllers.GetContracts)
	r.Post("contracts", controllers.AddContracts)

	r.Post("projects", controllers.AddProjects)
	r.Get("projects", controllers.GetProjects)
	r.Get("projects/:id", controllers.GetProjectsByID)

	r.Get("vm-specifications", controllers.GetAllVMSpecifications)
	r.Get("vm-specifications/:id", controllers.GetVMSpecificationsByID)
	r.Get("vm-specifications-by-projects/:id", controllers.GetVMSpecificationsByProjectID)
	r.Post("vm-specifications", controllers.AddVMSpecifications)

	r.Get("project-revenue", controllers.GetProjectRevenue)
	r.Get("total-revenue", controllers.GetTotalRevenue)
	r.Get("active-running-revenue", controllers.GetActiveRunningRevenue)
	r.Get("active-stopped-revenue", controllers.GetActiveStoppedRevenue)
	r.Get("highest-component", controllers.GetHighestResourceUsage)
	r.Get("segment-revenue", controllers.GetSegmentRevenue)
	r.Get("location-count", controllers.GetLocationCount)
}
