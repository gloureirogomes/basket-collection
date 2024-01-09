package domain

type Player struct {
	PlayerID string
	Name     string
	Age      int32
	Position string
	Number   int32
	Team     Team
}

func NewPlayer(playerID string, name string, age int32, position string, number int32, team Team) Player {
	return Player{
		PlayerID: playerID,
		Name:     name,
		Age:      age,
		Position: position,
		Number:   number,
		Team:     team,
	}
}

func (p Player) GetPlayerID() string {
	return p.PlayerID
}

func (p Player) GetName() string {
	return p.Name
}

func (p Player) GetAge() int32 {
	return p.Age
}

func (p Player) GetPosition() string {
	return p.Position
}

func (p Player) GetNumber() int32 {
	return p.Number
}

func (p Player) GetTeam() Team {
	return p.Team
}
