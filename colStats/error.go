package main

import "errors"

var (
	ErrNotNumber        = errors.New("Data is not a numeric")
	ErrInvalidColumn    = errors.New("Invalid column number")
	ErrNoFiles          = errors.New("No input files")
	ErrInvalidOperation = errors.New("Invalid Operation")
)
