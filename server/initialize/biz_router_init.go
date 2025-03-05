package initialize

import (
	"neocex/v2/router"

	"github.com/gofiber/fiber/v2"
)

func holder(routers ...fiber.Router) {
	_ = routers
	_ = router.RouterGroupApp
}

func initializeBizRouter(routers ...fiber.Router) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	privGroupFrontUser := routers[2]
	holder(publicGroup, privateGroup)
	{
		contractRouter := router.RouterGroupApp.Contract
		contractRouter.InitiateContractOrderRouter(privateGroup, publicGroup, privGroupFrontUser)
	}
}
