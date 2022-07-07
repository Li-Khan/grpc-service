package helper

import "errors"

var (
	ErrAlreadyExist  error = errors.New("name or date already exist")
	ErrEventNotFound error = errors.New("event not found")
	ErrInvalidDate   error = errors.New("invalid date")
)
