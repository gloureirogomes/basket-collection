package repository

import "github.com/GabrielLoureiroGomes/basket-collection/core/domain"

type TeamMongoDocument struct {
	Name       string `bson:"name"`
	Conference string `bson:"conference"`
	State      string `bson:"state"`
}

func newTeamMongoDocument(team *domain.Team) TeamMongoDocument {
	return TeamMongoDocument{
		Name:       team.GetName(),
		Conference: team.GetConference(),
		State:      team.GetState(),
	}
}

func newTeamListByTeamMongoDocument(teams []TeamMongoDocument) []*domain.Team {
	teamsToReturn := []*domain.Team{}

	for _, team := range teams {
		teamsToReturn = append(teamsToReturn, &domain.Team{
			Name:       team.Name,
			Conference: team.Conference,
			State:      team.State,
		})
	}

	return teamsToReturn
}
