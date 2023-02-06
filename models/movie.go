package models

type Movie struct {
	Name      string `json:"name" bson:"user_name"`
	Link      string `json:"link" bson:"link"`
	Duration  string `json:"duration" bson:"duration"`
	Publisher string `json:"publisher" bson:"publisher"`
}
