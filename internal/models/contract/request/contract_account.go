package request

type ChangeMarginType int

const(
	TransferIn ChangeMarginType = 1
	TransferOut ChangeMarginType = 2
)

type ChangeMarginReq struct {

}