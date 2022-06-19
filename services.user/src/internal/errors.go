package user

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrPasswordIncorrect  = errors.New("password incorrect")
	ErrUserExist          = errors.New("user already exist")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
