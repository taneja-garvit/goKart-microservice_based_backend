package models

type Order struct {
	ID string `json:"id" bson:"id"`
	UserID string `json:"user_id" bson:"user_id"`
	ItemName string `json:"item" bson:"item"`
}