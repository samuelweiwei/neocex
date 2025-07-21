package contract

import (
	"errors"
	"neocex/v2/internal/models/contract"
	"neocex/v2/internal/models/contract/request"
	"time"

	"github.com/google/uuid"
)

type ContractOrderService struct{}

func (c *ContractOrderService) CreateContractOrder(orderReq *request.ContractOrderReq, contractAccount *contract.ContractAccount, userID uint) error {
	if orderReq.UserId == 0 {
		return errors.New("Zero userid, illegal.........")
	}
	if orderReq.OrderType != contract.MarketOrder && orderReq.OrderType != contract.LimitOrder {
		return errors.New("Illegal order type of contract order")
	}
	if orderReq.OperationType != contract.OpenLong && orderReq.OperationType != contract.OpenShort {
		return errors.New("Illegal operation type of contract order")
	}

	//Business logic goes here

	var dbContractOrder = &contract.ContractOrder{
		Uuid:           uuid.New(),
		OrderTime:      time.Now(),
		UserId:         orderReq.UserId,
		SymbolId:       orderReq.SymbolId,
		SymbolName:     orderReq.SymbolName,
		OpenPrice:      orderReq.OpenPrice,
		Margin:         orderReq.Margin,
		LeverageRation: orderReq.LeverageRation,
		OrderType:      orderReq.OrderType,
		OperationType:  orderReq.OperationType,
		Status:         1,
	}
	_ = dbContractOrder
	return nil
}
