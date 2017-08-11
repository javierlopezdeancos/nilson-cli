package domain

type Championship struct {
	Name  string
	Teams []string
}

type Team struct {
	Name  string
	Alias string
}

type Player struct {
	Name  string
	Alias string
}

type Match struct {
	LocalTeam    string
	VisitorTeam  string
	LocalScore   int
	VisitorScore int
}
