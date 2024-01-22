package game

import (
	"example.com/quake/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGameResult(t *testing.T) {
	testCase := []Test{{
		Name: "ShouldDisplayGameInfo",
		Game: &model.Game{ID: 1, TotalKills: 11, Players: []model.Player{{ID: 1, Name: "Player 1", Kills: 11}}},
		Want: "\"game_1\":{\n  \"total_kills\": 11,\n  \"players\": [\n    \"Player 1\"\n  ],\n  \"kills\": {\n    \"Player 1\": 11\n  }\n}",
	}}

	for _, test := range testCase {
		t.Run(test.Name, func(t *testing.T) {
			resp, err := GetGameResult(*test.Game)
			assert.Nil(t, err)
			assert.Equal(t, test.Want, resp)
		})
	}
}

func TestGetGameWeaponsResult(t *testing.T) {
	testCase := []Test{{
		Name: "ShouldDisplayKillByMeansInfo",
		Game: &model.Game{ID: 1, TotalKills: 1, KillsByMeans: map[string]int{"MOD_TRIGGER_HURT": 1}},
		Want: "\"game_1\":{\n  \"kills_by_means\": {\n    \"MOD_TRIGGER_HURT\": 1\n  }\n}",
	}}

	for _, test := range testCase {
		t.Run(test.Name, func(t *testing.T) {
			resp, err := GetGameWeaponsResult(*test.Game)
			assert.Nil(t, err)
			assert.Equal(t, test.Want, resp)
		})
	}
}
