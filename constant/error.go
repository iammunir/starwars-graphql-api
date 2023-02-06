package constant

import "errors"

type errorKeyType int

const ErrorKey errorKeyType = iota

var (
	ErrCastingValue = errors.New("error casting value type")
	ErrorAuth       = errors.New("unauthorized")
	ErrorTimeout    = errors.New("processing time is too long")
	ErrorNotFound   = errors.New("record not found")
)
