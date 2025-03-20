package contract

import (
	"time"

	"github.com/ericlagergren/decimal"
	"github.com/google/uuid"
)

type OrderType int
type OperationType int
type OrderStatus int

// Market price - 1, Limited Price - 2, Clear Margin - 3
const (
	MarketOrder OrderType = 1
	LimitOrder  OrderType = 2
	ClearOrder  OrderType = 3
)

const (
	OpenLong   OperationType = 1
	OpenShort  OperationType = 2
	CloseLong  OperationType = 3
	CloseShort OperationType = 4
)

const (
	Pending         OrderStatus = 1
	PartiallyFilled OrderStatus = 2
	FullFilled      OrderStatus = 3
	Cancelled       OrderStatus = 4
)

type ContractOrder struct {
	Uuid           uuid.UUID     `json:"uuid" gorm:"primary_key" gorm:"type:uuid;default:uuid_generate_v4()"`
	OrderTime      time.Time     `json:"order_time" gorm:"comment:OrderTime"`
	UserId         uint          `json:"user_id" gorm:"comment:UserID"`
	SymbolId       uint          `json:"symbol_id" gorm:"comment:SymbolID"`
	SymbolName     string        `json:"symbol_name" gorm:"comment:SymbolName"`
	OpenPrice      decimal.Big   `json:"price" gorm:"comment:Price"`
	ClosePrice     decimal.Big   `json:"total_price" gorm:"comment:TotalPrice"`
	Margin         decimal.Big   `json:"margin" gorm:"comment:Margin"`
	LeverageRation int           `json:"leverage_ratio" gorm:"comment:LeverageRatio"`
	OrderType      OrderType     `json:"order_type" gorm:"comment:OrderType"`
	OperationType  OperationType `json:"operation_type" gorm:"comment:OperationType"`
	Status         OrderStatus   `json:"status" gorm:"comment:Status"`
}

func (ContractOrder) TableName() string {
	return "contract_order"
}
