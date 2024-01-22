package game

import (
	"testing"

	"example.com/quake/internal/model"
	"example.com/quake/internal/utils"
	"github.com/stretchr/testify/assert"
)

type Test struct {
	Name   string
	Line   string
	Player *model.Player
	Game   *model.Game
	Want   interface{}
	Error  error
	Index  *int
}

func intToPointer(i int) *int {
	return &i
}

func TestNewPlayer(t *testing.T) {
	players := []model.Player{{ID: 0, Name: "Player 0"}, {ID: 1, Name: "Player 1"}}

	testCases := []Test{{
		Name:   "ShouldReturnAsNewPlayerJoining",
		Player: &model.Player{ID: 2, Name: "Player 2"},
		Want:   true,
		Index:  nil,
	}, {
		Name:   "ShouldReturnAsUpdatingPlayerInfo",
		Player: &model.Player{ID: 0, Name: "Player 0 name updated"},
		Want:   false,
		Index:  intToPointer(0),
	}}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			found, index := isNewPlayer(test.Player, players)
			assert.Equal(t, found, test.Want)
			assert.Equal(t, index, test.Index)
		})
	}
}

func TestParsePlayer(t *testing.T) {
	testCases := []Test{{
		Name: "ShouldReturnPlayerInfo",
		Line: "21:51 ClientUserinfoChanged: 3 n\\Dono da Bola\\t\\0\\model\\sarge/krusade\\hmodel\\sarge/krusade\\g_redteam\\\\g_blueteam\\\\c1\\5\\c2\\5\\hc\\95\\w\\0\\l\\0\\tt\\0\\tl\\0",
		Want: &model.Player{ID: 3, Name: "Dono da Bola"},
	}, {
		Name: "ShouldReturnNilMatchesNotFound",
		Line: "wrong line to be read",
		Want: (*model.Player)(nil),
	}}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			resp := parsePlayer(test.Line)
			assert.Equal(t, test.Want, resp)
		})
	}
}

func TestParseKill(t *testing.T) {
	testCases := []Test{{
		Name: "ShouldReturnRemoveKillCount",
		Line: "21:07 Kill: 1022 2 22: <world> killed Player 2 by MOD_TRIGGER_HURT",
		Game: &model.Game{
			ID:           1,
			Players:      []model.Player{{ID: 1, Name: "Player 1"}, {ID: 2, Name: "Player 2"}},
			KillsByMeans: make(map[string]int),
		},
		Want: &model.Game{
			ID:           1,
			TotalKills:   1,
			Players:      []model.Player{{ID: 1, Name: "Player 1", Kills: 0}, {ID: 2, Name: "Player 2", Kills: -1}},
			KillsByMeans: map[string]int{"MOD_TRIGGER_HURT": 1},
		},
	}, {
		Name: "ShouldReturnAddingKillCount",
		Line: "21:07 Kill: 1 2 22: Player 1 killed Player 2 by MOD_TRIGGER_HURT",
		Game: &model.Game{
			ID:           1,
			Players:      []model.Player{{ID: 1, Name: "Player 1"}, {ID: 2, Name: "Player 2"}},
			KillsByMeans: make(map[string]int),
		},
		Want: &model.Game{
			ID:           1,
			TotalKills:   1,
			Players:      []model.Player{{ID: 1, Name: "Player 1", Kills: 1}, {ID: 2, Name: "Player 2", Kills: 0}},
			KillsByMeans: map[string]int{"MOD_TRIGGER_HURT": 1},
		},
	}, {
		Name: "ShouldReturnErrorPlayerNotFound",
		Line: "21:07 Kill: 1 2 22: Player 1 killed Player 2 by MOD_TRIGGER_HURT",
		Game: &model.Game{
			ID:           1,
			Players:      []model.Player{},
			KillsByMeans: make(map[string]int),
		},
		Want: &model.Game{
			ID:           1,
			TotalKills:   0,
			Players:      []model.Player{},
			KillsByMeans: map[string]int{},
		},
		Error: utils.PlayerNotFound,
	}, {
		Name:  "ShouldReturnErrorParsingKillLine",
		Line:  "",
		Want:  (*model.Game)(nil),
		Error: utils.ParseKillLine,
	}}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			resp := parseKill(test.Game, test.Line)
			assert.Equal(t, resp, test.Error)
			assert.Equal(t, test.Want, test.Game)
		})
	}
}

func TestParseLogFile(t *testing.T) {
	testCase := []Test{{
		Name: "ShouldReadFile",
		Line: "../test/testFile.log",
		Want: []model.Game{{
			ID:           1,
			TotalKills:   11,
			Players:      []model.Player{{ID: 2, Name: "Isgalamido", Kills: -5}, {ID: 3, Name: "Mocinha", Kills: 0}},
			KillsByMeans: map[string]int{"MOD_TRIGGER_HURT": 7, "MOD_ROCKET_SPLASH": 3, "MOD_FALLING": 1},
		}},
		Error: nil,
	}, {
		Name:  "ShouldReturnErrorFileNotFound",
		Line:  "../test/nonExistentFile",
		Want:  ([]model.Game)(nil),
		Error: utils.FileNotFound,
	}}

	for _, test := range testCase {
		t.Run(test.Name, func(t *testing.T) {
			resp, err := ParseLogFile(test.Line)
			assert.Equal(t, test.Error, err)
			assert.Equal(t, test.Want, resp)
		})
	}

}
