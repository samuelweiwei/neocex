package initialize

import "github.com/gofiber/fiber/v2"

func Routers() *fiber.App {
	Router := fiber.New()
	PublicGroup := Router.Group("/contractorder")
	PrivateGroup := Router.Group("/contractorder")
	PrivateGroupFrontUser := Router.Group("/contractorder")
	initializeBizRouter(PrivateGroup, &PublicGroup, &PrivateGroupFrontUser)
	return Router
}
