// add areas types
package types

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

type Items struct {
	ID         int    `json:"id"`
	ItemName   string `json:"item_name"`
	ItemType   string `json:"item_type"`
	ItemRarity string `json:"item_rarity"`
	ItemLevel  int    `json:"item_level"`
	ItemValue  int    `json:"item_value"`
}

type Stats struct {
	ID           int `json:"id"`
	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Intelligence int `json:"intelligence"`
	Wisdom       int `json:"wisdom"`
	Charisma     int `json:"charisma"`
}

type Areas struct {
	ID        int    `json:"id"`
	AreaName  string `json:"area_name"`
	AreaLevel int    `json:"area_level"`
	AreaType  string `json:"area_type"`
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