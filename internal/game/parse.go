package game

import (
	"bufio"
	"example.com/quake/internal/model"
	"example.com/quake/internal/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ParseLogFile reads log file and extract its data
func ParseLogFile(filePath string) ([]model.Game, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, utils.FileNotFound
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	games := make([]model.Game, 0)
	var currentGame *model.Game

	for scanner.Scan() {
		line := scanner.Text()

		// Checks if it's a new game
		if strings.Contains(line, utils.GameInt) {
			if currentGame != nil {
				games = append(games, *currentGame)
			}
			currentGame = &model.Game{ID: len(games) + 1, KillsByMeans: make(map[string]int)}
		}

		// Check player updates (name)
		if strings.Contains(line, utils.PlayerData) {
			player := parsePlayer(line)
			if player != nil {
				if isNew, index := isNewPlayer(player, currentGame.Players); isNew {
					currentGame.Players = append(currentGame.Players, *player)
				} else {
					currentGame.Players[*index].Name = player.Name
				}
			}
		}

		// Get kill info
		if strings.Contains(line, utils.Kill) {
			err := parseKill(currentGame, line)
			if err != nil {
				return nil, err
			}
		}
	}

	if currentGame != nil {
		games = append(games, *currentGame)
	}

	return games, nil
}

// parseKill extract kills info and update game's object
func parseKill(currentGame *model.Game, line string) error {
	regex := regexp.MustCompile(`Kill:\s+(\d+)\s+(\d+)\s+(\d+): (.+?) killed (.+?) by (.+)`)
	matches := regex.FindStringSubmatch(line)

	if len(matches) == 0 {
		return utils.ParseKillLine
	}

	killerID, err := strconv.Atoi(matches[1])
	if err != nil {
		return err
	}

	victimID, err := strconv.Atoi(matches[2])
	if err != nil {
		return err
	}

	killModeID, err := strconv.Atoi(matches[3])
	if err != nil {
		return err
	}

	found := false
	for i, player := range currentGame.Players {
		if killerID == utils.World && victimID == player.ID {
			currentGame.Players[i].Kills -= 1
			found = true
			break
		} else {
			if killerID == player.ID {
				currentGame.Players[i].Kills += 1
				found = true
				break
			}
		}
	}

	if !found {
		return utils.PlayerNotFound
	}

	if mode, ok := utils.KillMode[killModeID]; ok {
		currentGame.KillsByMeans[mode]++
		currentGame.TotalKills++
	}

	return nil
}

// parsePlayer extract info about a player
func parsePlayer(line string) *model.Player {
	regex := regexp.MustCompile(`ClientUserinfoChanged:\s+(\d+)\s+n\\(.+?)\\t\\`)
	matches := regex.FindStringSubmatch(line)

	if len(matches) == 0 {
		return nil
	}

	playerID, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil
	}

	playerName := matches[2]

	return &model.Player{
		ID:   playerID,
		Name: playerName,
	}
}

// isNewPlayer check if it's a new player if not returns its index
func isNewPlayer(newPlayer *model.Player, players []model.Player) (bool, *int) {
	for i, player := range players {
		if player.ID == newPlayer.ID {
			return false, &i
		}
	}
	return true, nil
}
