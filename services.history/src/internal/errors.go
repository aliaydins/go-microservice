package history

import "errors"

var (
	ErrNotValidID      = errors.New("id param is not valid")
	ErrNotFoundAnyData = errors.New("not found any data with id")
)
