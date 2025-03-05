package router

import "neocex/v2/router/contract"

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	Contract contract.RouteGroup
}
