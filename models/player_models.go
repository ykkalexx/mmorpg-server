package models

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