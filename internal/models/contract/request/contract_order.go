package request

import (
	"neocex/v2/internal/models/contract"

	"github.com/ericlagergren/decimal"
)

type ContractOrderReq struct {
	Uuid           string                 `json:"uuid"`
	UserId         uint                   `json:"user_id"`
	SymbolId       uint                   `json:"symbol_id"`
	SymbolName     string                 `json:"symbol_name"`
	OpenPrice      decimal.Big            `json:"price"`
	OperationType  contract.OperationType `json:"total_price"`
	OrderType      contract.OrderType     `json:"order_type"`
	Margin         decimal.Big            `json:"margin"`
	LeverageRation int                    `json:"leverage_ratio"`
}
