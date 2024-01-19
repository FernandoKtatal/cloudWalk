package game

import (
	"encoding/json"
	"example.com/quake/internal/model"
	"example.com/quake/internal/utils"
)

// GetGameResult returns data from each match
func GetGameResult(game model.Game) (string, error) {
	out := &model.GameResponse{Kills: make(map[string]int)}
	out.TotalKills = game.TotalKills

	for _, player := range game.Players {

		out.Players = append(out.Players, player.Name)
		out.Kills[player.Name] = player.Kills
	}

	resp, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return "", err
	}

	return utils.GameKeyWithID(game.ID, string(resp)), nil

}

func GetGameWeaponsResult(game model.Game) (string, error) {
	out := &model.GameWeaponsResponse{KillsByMeans: make(map[string]int)}

	out.KillsByMeans = game.KillsByMeans
	resp, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return "", err
	}

	return utils.GameKeyWithID(game.ID, string(resp)), nil
}
