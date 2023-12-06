package domain

type Player struct {
	PlayerID string
	Name     string
	Age      int32
	Position string
	Country  string
	Team     Team
	Height   float32
	Weight   int32
}

func NewPlayer(playerID string, name string, age int32, position string, country string, team Team, height float32, weight int32) Player {
	return Player{
		PlayerID: playerID,
		Name:     name,
		Age:      age,
		Position: position,
		Country:  country,
		Team:     team,
		Height:   height,
		Weight:   weight,
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

func (p Player) GetCountry() string {
	return p.Country
}

func (p Player) GetTeam() Team {
	return p.Team
}

func (p Player) GetHeight() float32 {
	return p.Height
}

func (p Player) GetWeight() int32 {
	return p.Weight
}
