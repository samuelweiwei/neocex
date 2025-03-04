package contract

import service "neocex/v2/internal/service"

type ApiGroup struct {
	ContractOrderAggregate
}

var (
	contractOrderServ = service.ServiceGroupApp.ContractServiceGroup.ContractOrderService
)
