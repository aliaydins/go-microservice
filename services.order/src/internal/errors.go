package order

import "errors"

var (
	ErrOrdersNotFound   = errors.New("orders not found")
	ErrOrderDeny        = errors.New("you have not enough balance for the order")
	ErrOrderNotCreated  = errors.New("error occured when order saving to database ")
	ErrWhenOrderDeleted = errors.New("error occured when order deleted")
	ErrWalletUpdated    = errors.New("error occured when updating wallet")
)
