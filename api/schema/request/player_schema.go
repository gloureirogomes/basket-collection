package request

import (
	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
)

// InsertPlayerSchema is a struct to bind parameter for insert player data
type InsertPlayerSchema struct {
	Name     string      `json:"name"`
	Age      int32       `json:"age"`
	Position string      `json:"position"`
	Country  string      `json:"country"`
	Team     domain.Team `json:"team"`
	Height   float32     `json:"height"`
	Weight   int32       `json:"weight"`
}
