package es

import "errors"

var (
	ErrAlreadyExists       = errors.New("already exists")
	ErrAggregateNotFound   = errors.New("aggregate not found")
	ErrInvalidEventType    = errors.New("invalid event type")
	ErrInvalidCommandType  = errors.New("invalid command type")
	ErrInvalidAggregate    = errors.New("invalid aggregate")
	ErrInvalidAggregateID  = errors.New("invalid aggregate id")
	ErrInvalidEventVersion = errors.New("invalid event version")
)
