package wallet

import "errors"

var (
	ErrWalletNotFound   = errors.New("wallet not found")
	ErrNotValidWalletID = errors.New("id param is not valid")
	ErrNotValidUserID   = errors.New("user id  is not valid")
	ErrWalletNotCreated = errors.New("error occured when created user wallet")
	ErrWalletNotUpdated = errors.New("error occured when updated wallet")
)
