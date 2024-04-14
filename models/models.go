package models

type LoginDTO struct {
	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`
}

type RegisterDTO struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}
