package models

type Order struct {
	ID string `json:"id"`
	UserID string `json:"user_id"`
	ItemName string `json:"item_name"`
}