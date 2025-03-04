package contract

import (
	"neocex/v2/i18n"
	contractOrderReq "neocex/v2/internal/models/contract/request"
	"neocex/v2/internal/service/contract"
	"neocex/v2/logging"
	"neocex/v2/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ContractOrderAggregate struct{}

var contractAccountService = contract.ContractAccountService{}
var contractOrderService = contract.ContractOrderService{}

func (c *ContractOrderAggregate) CreateContractOrder(f *fiber.Ctx) error {
	var contractOrder contractOrderReq.ContractOrderReq
	if err := f.BodyParser(&contractOrder); err != nil {
		logging.Logger.Error(i18n.T("error.parsing_request_body", nil), zap.Error(err))
		return err
	}
	userID := utils.GetUserID(f)
	contractAccount, err := contractAccountService.GetContractAccount(userID)
	if err != nil {
		logging.Logger.Error(i18n.T("error.getting_contract_account", nil), zap.Error(err))
		return err
	}
	err = contractOrderService.CreateContractOrder(&contractOrder, &contractAccount, userID)
	if err != nil {
		logging.Logger.Error(i18n.T("error.creating_contract_order", nil), zap.Error(err))
	}
	return nil
}
func (c *ContractOrderAggregate) UpdateContractOrder() {}
func (c *ContractOrderAggregate) DeleteContractOrder() {}
func (c *ContractOrderAggregate) GetContractOrder()    {}
