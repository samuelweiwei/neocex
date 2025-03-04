package contract

import (
	"neocex/v2/internal/models/contract"
	"neocex/v2/internal/models/contract/request"
)

type ContractOrderService struct{}

func (c *ContractOrderService) CreateContractOrder(orderReq *request.ContractOrderReq, contractAccount *contract.ContractAccount, userID uint) error {
	return nil
}
