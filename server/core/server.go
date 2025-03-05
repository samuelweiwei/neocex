package core

import "neocex/v2/server/initialize"

func StartWindowsServer() {
	Router := initialize.Routers()
	_ = Router
}
