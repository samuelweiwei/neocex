package sercice

import "neocex/v2/internal/service/contract"

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	ContractServiceGroup contract.ServiceGroup
}
