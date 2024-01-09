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

func newPlayerListByPlayerMongoDocument(players []PlayerMongoDocument) []domain.Player {
	playersToReturn := []domain.Player{}

	for _, player := range players {
		playersToReturn = append(playersToReturn, domain.Player{
			Name:     player.Name,
			Age:      player.Age,
			Position: player.Position,
			Number:   player.Number,
			Team:     player.Team,
		})
	}

	return playersToReturn
}
