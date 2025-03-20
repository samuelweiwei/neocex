package symbol

import (
	"context"
	"fmt"
	"neocex/v2/global"
	"strconv"
)

type SymbolService struct{}

func (ss *SymbolService) GetSymbolPriceById(symbolId string) (interface{}, error) {
	var symbolRecord struct {
		Symbol       string
		Type         *int
		CurrentPrice *float64
	}

	if err := global.GVA_DB.Table("symbols").Select("symbol", "type", "current_price").
		Where("id = ?", symbolId).First(&symbolRecord).Error; err != nil {
		return nil, fmt.Errorf("Symbol does not exist...")
	}
	if symbolRecord.Type == nil {
		return nil, fmt.Errorf("Unknown symbol type...")
	}

	//Construct the redis key based on type
	var redisKey string
	switch *symbolRecord.Type {
	case 0: //Stock
		redisKey = fmt.Sprintf("symbol:stock:%s", symbolRecord.Symbol)
	case 1: //Crypto
		redisKey = fmt.Sprintf("symbol:crypto:%s", symbolRecord.Symbol)
	case 2: //Forex
		redisKey = fmt.Sprintf("symbol:forex:%s", symbolRecord.Symbol)
	default:
		return nil, fmt.Errorf("Unknown symbol type:%d", *symbolRecord.Type)
	}
	_ = redisKey

	if price, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result(); err == nil {
		priceFloat, _ := strconv.ParseFloat(price, 64)
		return priceFloat, nil
	}

	if symbolRecord.CurrentPrice != nil {
		return *symbolRecord.CurrentPrice, nil
	}

	return nil, fmt.Errorf("Symbol price not found...")
}
