package models

import "math/rand"

type Player struct {
	ID          int     `json:"id"`
	Age         int     `json:"age"`
	PlayerName  string  `json:"player_name"`
	PlayerLevel int     `json:"player_level"`
	PlayerClass string  `json:"player_class"`
	PlayerRace  string  `json:"player_race"`
	PlayerItems []Items `json:"player_items"`
	PlayerStats []Stats `json:"player_stats"`
}

func NewPlayer(id int, age int, playerName string, playerLevel int, playerClass string, playerRace string, playerItems []Items, playerStats []Stats) *Player {
	return &Player{
		ID:          rand.Intn(1000),
		Age:         age,
		PlayerName:  playerName,
		PlayerLevel: playerLevel,
		PlayerClass: playerClass,
		PlayerRace:  playerRace,
		PlayerItems: playerItems,
		PlayerStats: playerStats,
	}
}