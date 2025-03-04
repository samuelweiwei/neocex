package contract

import (
	"time"

	"github.com/ericlagergren/decimal"
)

type AccountStatus int

const (
	Unreviewed AccountStatus = 1
	Reviewed   AccountStatus = 2
)

type ContractAccount struct {
	UserId             uint          `json:"userId" form:"userId" gorm:"column:user_id;comment:Relate to user name ID;size:20;"`
	TotalMargin        decimal.Big   `json:"totalMargin" form:"totalMargin" gorm:"column:total_margin;comment:Total margin in account;size:20;"`
	AvailableMargin    decimal.Big   `json:"availableMargin" form:"availableMargin" gorm:"column:available_margin;comment:Applicable margin in account;size:20;"`
	FrozenMargin       decimal.Big   `json:"frozenMargin" form:"frozenMargin" gorm:"column:frozen_margin;comment:Frozen margin in account;size:20;"`
	UsedMargin         decimal.Big   `json:"usedMargin" form:"usedMargin" gorm:"column:used_margin;comment:Used margin in account;size:20;"`
	RealizedProfitLoss decimal.Big   `json:"realizedProfitLoss" form:"realizedProfitLoss" gorm:"column:realized_profit_loss;comment:Realized profit loss in account;size:20;"`
	AccountStatus      AccountStatus `json:"accountStatus" form:"accountStatus" gorm:"column:account_status;comment:Account status;size:20;"`
	CreatedBy          uint          `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:Created by;size:20;"`
	UpdatedBy          uint          `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:Updated by;size:20;"`
	DeletedBy          uint          `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:Deleted by;size:20;"`
	CreatedAt          time.Time     `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:Created time;"`
}
