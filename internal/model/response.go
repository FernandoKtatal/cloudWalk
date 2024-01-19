package model

type GameResponse struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

type GameWeaponsResponse struct {
	KillsByMeans map[string]int `json:"kills_by_means"`
}
