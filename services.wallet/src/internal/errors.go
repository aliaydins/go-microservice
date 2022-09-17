package wallet

import "errors"

var (
	ErrWalletNotFound   = errors.New("wallet not found")
	ErrNotValidWalletID = errors.New("id param is not valid")
	ErrWalletNotCreated = errors.New("error occured when created user wallet")
)
