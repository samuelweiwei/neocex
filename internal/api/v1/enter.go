package v1

import "neocex/v2/internal/api/v1/contract"

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	ContractApiGroup contract.ApiGroup
}
