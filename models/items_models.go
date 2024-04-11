package models

type Items struct {
	ID         int    `json:"id"`
	ItemName   string `json:"item_name"`
	ItemType   string `json:"item_type"`
	ItemRarity string `json:"item_rarity"`
	ItemLevel  int    `json:"item_level"`
	ItemValue  int    `json:"item_value"`
}