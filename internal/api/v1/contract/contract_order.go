package contract

import (
	"neocex/v2/i18n"
	contractOrderReq "neocex/v2/internal/models/contract/request"
	"neocex/v2/logging"
	"neocex/v2/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ContractOrderAggregate struct{}

func (c *ContractOrderAggregate) CreateContractOrder(f *fiber.Ctx) error {
	var contractOrder contractOrderReq.ContractOrder
	if err := f.BodyParser(&contractOrder); err != nil {
		logging.Logger.Error(i18n.T("error.parsing_request_body", nil), zap.Error(err))
		return err
	}
	userID := utils.GetUserID(f)
	return nil
}
func (c *ContractOrderAggregate) UpdateContractOrder() {}
func (c *ContractOrderAggregate) DeleteContractOrder() {}
func (c *ContractOrderAggregate) GetContractOrder()    {}
