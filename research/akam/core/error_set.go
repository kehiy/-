package core

import "errors"

var (
	ErrorGet     = errors.New("key not found")
	ErrorExecute = errors.New("invalid command")
)
