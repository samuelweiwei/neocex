package contract

import (
	"time"

	"github.com/ericlagergren/decimal"
)

type PositionType int
type PositionStatus int

const (
	Long  PositionType = 1
	Short PositionType = 2
)

const (
	Unclosed PositionStatus = 1
	Closed   PositionStatus = 2
)

type ContractPosition struct {
	UserId          uint
	SymbolId        uint
	SymbolName      string
	PositionTime    time.Time
	Quantity        float64
	OpenPrice       decimal.Big
	LeverageRatio   int
	Margin          decimal.Big
	PositionAmount  float64
	ForceClosePrice decimal.Big
	PositionType    PositionType
	PositionStatus  PositionStatus
	CreatedBy       uint
	UpdatedBy       uint
	DeletedBy       uint
}
