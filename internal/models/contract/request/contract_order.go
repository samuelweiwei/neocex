package request

type ContractOrderReq struct {
	Uuid           string  `json:"uuid"`
	UserId         uint    `json:"user_id"`
	SymbolId       uint    `json:"symbol_id"`
	SymbolName     string  `json:"symbol_name"`
	Price          float64 `json:"price"`
	TotalPrice     float64 `json:"total_price"`
	Margin         float64 `json:"margin"`
	LeverageRation float64 `json:"leverage_ratio"`
	Status         uint    `json:"status"`
}
