package wallet

import "errors"

var (
	ErrWalletNotFound   = errors.New("wallet not found")
	ErrNotValidWalletID = errors.New("id param is not valid")
)
