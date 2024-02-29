package errapp

import (
	"errors"
)

var (
	ErrorNotFoundUser    = errors.New("not found user")
	ErrorNotFoundProduct = errors.New("not found product")
)
