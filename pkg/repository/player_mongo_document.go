package repository

import "github.com/GabrielLoureiroGomes/basket-collection/core/domain"

type PlayerMongoDocument struct {
	Name     string      `bson:"name"`
	Age      int32       `bson:"age"`
	Position string      `bson:"position"`
	Number   int32       `bson:"number"`
	Team     domain.Team `bson:"team"`
}

func newPlayerMongoDocument(player domain.Player) PlayerMongoDocument {
	return PlayerMongoDocument{
		Name:     player.GetName(),
		Age:      player.GetAge(),
		Position: player.GetPosition(),
		Number:   player.GetNumber(),
		Team:     player.GetTeam(),
	}
}
