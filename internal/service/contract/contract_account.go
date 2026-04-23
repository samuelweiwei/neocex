package contract

import (
	"neocex/v2/internal/models/contract"
)

type ContractAccountService struct{}

func (contractAccountService *ContractAccountService) GetContractAccount(userID uint) (contractAccount contract.ContractAccount, err error) {
	//hard coding for now
	return contractAccount, nil
}
