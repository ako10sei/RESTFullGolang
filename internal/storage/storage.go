package storage

import "errors"

var (
	ErrURLNotFound       = errors.New("url not found")
	ErrURLAlreadyExists  = errors.New("url already exists")
	ErrURLIsNotSupported = errors.New("url is not supported")
	ErrURLIsInvalid      = errors.New("url is invalid")
)
