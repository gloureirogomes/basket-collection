package domain

type Team struct {
	TeamID     string
	Name       string
	Conference string
	State      string
}

func NewTeam(teamID string, name string, conference string, state string) Team {
	return Team{
		TeamID:     teamID,
		Name:       name,
		Conference: conference,
		State:      state,
	}
}

func (t Team) GetTeamID() string {
	return t.TeamID
}

func (t Team) GetName() string {
	return t.Name
}

func (t Team) GetConference() string {
	return t.Conference
}

func (t Team) GetState() string {
	return t.State
}
