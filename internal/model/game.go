package model

// Player represents a player inside the game
type Player struct {
	ID    int
	Name  string
	Kills int
}

// Game represents a match
type Game struct {
	ID           int            `json:"ID"`
	TotalKills   int            `json:"total_kills"`
	Players      []Player       `json:"players"`
	KillsByMeans map[string]int `json:"kills_by_means"`
}
