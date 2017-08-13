package data

import (
	"github.com/google/uuid"
)

// https://data-devnilson.wedeploy.io/championships
type Championship struct {
	id uuid.UUID
	Name  string
	Teams []uuid.UUID
}

// https://data-devnilson.wedeploy.io/teams
type Team struct {
	id uuid.UUID
	Name  string
	Alias string
}

// https://data-devnilson.wedeploy.io/players
type Player struct {
	id uuid.UUID
	Name  string
	Alias string
}

// https://data-devnilson.wedeploy.io/matches
type Match struct {
	id uuid.UUID
	LocalTeam    string
	VisitorTeam  string
	LocalScore   int
	VisitorScore int
}
