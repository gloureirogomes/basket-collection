package request

// InsertPlayerSchema is a struct to bind parameter for insert player data
type InsertPlayerSchema struct {
	Name       string `json:"name"`
	Age        int32  `json:"age"`
	Position   string `json:"position"`
	Number     int32  `json:"number"`
	TeamName   string `json:"team_name"`
	Conference string `json:"conference"`
}
