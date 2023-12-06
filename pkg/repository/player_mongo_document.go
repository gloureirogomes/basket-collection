package repository

import "github.com/GabrielLoureiroGomes/basket-collection/core/domain"

type PlayerMongoDocument struct {
	Name     string      `bson:"name"`
	Age      int32       `bson:"age"`
	Position string      `bson:"position"`
	Country  string      `bson:"country"`
	Team     domain.Team `bson:"team"`
	Height   float32     `bson:"height"`
	Weight   int32       `bson:"weight"`
}

func newPlayerMongoDocument(player domain.Player) PlayerMongoDocument {
	return PlayerMongoDocument{
		Name:     player.GetName(),
		Age:      player.GetAge(),
		Position: player.GetPosition(),
		Country:  player.GetCountry(),
		Team:     player.GetTeam(),
		Height:   player.GetHeight(),
		Weight:   player.GetWeight(),
	}
}
