package model

import "errors"

var (
	// ErrPersonCanNotBeNil la persona no puede ser nula
	ErrPersonCanNotBeNil = errors.New("la persona no puede ser nula")
	// ErrIDPersonDoesNotExists el id de la persona no existe
	ErrIDPersonDoesNotExists = errors.New("la persona no existe")
)
