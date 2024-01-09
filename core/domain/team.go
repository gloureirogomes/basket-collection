package domain

type Team struct {
	Name       string
	Conference string
}

func NewTeam(name string, conference string) Team {
	return Team{
		Name:       name,
		Conference: conference,
	}
}

func (t Team) GetName() string {
	return t.Name
}

func (t Team) GetConference() string {
	return t.Conference
}
