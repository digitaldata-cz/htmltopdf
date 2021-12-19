package htmltopdf

import "errors"

var (
	ErrLibratyNotInitialized  = errors.New("library not initialized")
	ErrCreateGlobalSettings   = errors.New("can not create global settings")
	ErrCreateConverter        = errors.New("can not create converter")
	ErrUninitializedConverter = errors.New("uninitialized converter")
	ErrUninitializedWriter    = errors.New("uninitialized writer")
	ErrNoObjects              = errors.New("no objects")
	ErrConversionFailed       = errors.New("conversion failed")
	ErrInvalidOption          = errors.New("invalid option")
)
