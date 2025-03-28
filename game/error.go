package game

import (
	"errors"
)

var (
	InvalidCoordinate = errors.New("invalid coordinate")
	InvalidMove       = errors.New("invalid move")
)
