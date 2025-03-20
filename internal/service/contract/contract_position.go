package contract

import (
	"fmt"
	"neocex/v2/global"
	"neocex/v2/internal/models/contract"
	"neocex/v2/internal/service/symbol"

	"github.com/ericlagergren/decimal"
)

type ContractPositionService struct{}

func (c *ContractPositionService) GetUnrealizedProfitLoss(userId uint) (unrealizedProfitLoss decimal.Big, err error) {
	var positions []contract.ContractPosition
	err = global.GVA_DB.Model(&contract.ContractPosition{}).Where("user_id = ? AND position_status = ?", userId, contract.Unclosed).Find(&positions).Error
	if err != nil {
		return
	}
	for _, position := range positions {
		_ = position
		_, single, err := c.GetUnrealizedProfitLossSingle(int(position.SymbolId), *new(decimal.Big).SetFloat64(position.Quantity), position.OpenPrice, position.PositionType)
		_ = single
		if err != nil {
			return *new(decimal.Big).SetMantScale(0, 0), err
		}

	}
}

func (c *ContractPositionService) GetUnrealizedProfitLossSingle(symbolId int, quantity decimal.Big, openPrice decimal.Big, positionType contract.PositionType) (closePrice decimal.Big, unrealizedProfitLoss decimal.Big, err error) {
	//Acquire the position profit and loss in redis generally
	symbolService := symbol.SymbolService{}
	symbolPrice, err := symbolService.GetSymbolPriceById(fmt.Sprintf("%d", symbolId))
	if err != nil {
		return *new(decimal.Big).SetMantScale(0, 0), *new(decimal.Big).SetMantScale(0, 0), err
	}

	closePrice = *new(decimal.Big).SetFloat64(symbolPrice.(float64))
	var finalPrice decimal.Big
	switch positionType {
	case contract.Long:
		unrealizedProfitLoss = *finalPrice.Sub(&openPrice, &closePrice).Mul(&finalPrice, &quantity).Round(2)
	case contract.Short:
		unrealizedProfitLoss = *finalPrice.Sub(&openPrice, &closePrice).Mul(&finalPrice, &quantity).Round(2)
	}
	return

}
