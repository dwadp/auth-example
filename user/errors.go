package user

import "errors"

var (
	NotFound      = errors.New("user cannot be found")
	WrongPassword = errors.New("wrong password for this user")
)
