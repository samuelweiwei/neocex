package contract

import service "neocex/v2/internal/service/contract"

type RouteGroup struct {
	ContractOrderRouter
}

var (
	contractOrderApi = service.ContractOrderService{}
)
