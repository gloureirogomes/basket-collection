package service

import (
	"context"
)

// Player interface is used to manage player functions
type Player interface {

	// InsertPlayer function is used to create a player register
	InsertPlayer(ctx context.Context) error
}
