package contract

import (
	"time"

	"github.com/google/uuid"
)

type ContractOrder struct {
	Uuid           uuid.UUID `json:"uuid" gorm:"primary_key" gorm:"type:uuid;default:uuid_generate_v4()"`
	OrderTime      time.Time `json:"order_time" gorm:"comment:OrderTime"`
	UserId         uint      `json:"user_id" gorm:"comment:UserID"`
	SymbolId       uint      `json:"symbol_id" gorm:"comment:SymbolID"`
	SymbolName     string    `json:"symbol_name" gorm:"comment:SymbolName"`
	Price          float64   `json:"price" gorm:"comment:Price"`
	TotalPrice     float64   `json:"total_price" gorm:"comment:TotalPrice"`
	Margin         float64   `json:"margin" gorm:"comment:Margin"`
	LeverageRation float64   `json:"leverage_ratio" gorm:"comment:LeverageRatio"`
	OrderType      uint      `json:"order_type" gorm:"comment:OrderType"`
	Status         uint      `json:"status" gorm:"comment:Status"`
}
