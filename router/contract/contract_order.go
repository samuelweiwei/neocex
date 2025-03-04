package contract

import (
	"github.com/gofiber/fiber/v2"
	api "neocex/v2/internal/api/v1"
)

type ContractOrderRouter struct{}

func (s *ContractOrderRouter) InitiateContractOrderRouter(router fiber.Router, publicRouter fiber.Router, frontUserRouter fiber.Router) {
	// Create route groups
	contractOrderRouter := router.Group("/contractOrder")
	// contractOrderRouterWithoutRecord := router.Group("/contractOrder")
	// contractOrderRouterWithoutAuth := publicRouter.Group("/contractOrder")
	// contractOrderRouterFrontUser := frontUserRouter.Group("/contractOrder")
	{
		contractOrderRouter.Post("createContractOrder", api.ApiGroupApp.ContractApiGroup.CreateContractOrder)
		// contractOrderRouter.Put("/", s.UpdateContractOrder)
		// contractOrderRouter.Delete("/", s.DeleteContractOrder)
		// contractOrderRouter.Get("/", s.GetContractOrder)
	}
}
