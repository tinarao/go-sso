package models

type LoginDTO struct {
	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`
}
