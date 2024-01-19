package model

// Player represents a player inside the game
type Player struct {
	ID    int
	Name  string
	Kills int
}

// Game represents a match
type Game struct {
	ID           int
	TotalKills   int
	Players      []Player
	KillsByMeans map[string]int
}
