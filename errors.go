package gobitcask

import "errors"

var (
	ErrKeyNotFound = errors.New("Key not found in database")
	ErrInvalidDBFile = errors.New("Invalid DB file")
)
