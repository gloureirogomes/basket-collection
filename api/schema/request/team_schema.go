package request

import (
	"fmt"
)

// InsertTeamSchema is a struct to bind parameter for insert team data
type InsertTeamSchema struct {
	Name       string `json:"name"`
	Conference string `json:"conference"`
	State      string `json:"state"`
}

// FindOneTeamSchema is a struct to bind parameter for get one team data
type FindOneTeamSchema struct {
	Name string `form:"name"`
}

func (t InsertTeamSchema) IsValid() error {
	if t.Name == "" {
		return fmt.Errorf("invalid team name")
	}

	if t.Conference == "" {
		return fmt.Errorf("invalid team conference")
	}

	if t.State == "" {
		return fmt.Errorf("invalid team state")
	}

	return nil
}

func (t FindOneTeamSchema) IsValid() error {
	if t.Name == "" {
		return fmt.Errorf("invalid team name")
	}

	return nil
}
