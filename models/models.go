package models

import "time"

type LoginDTO struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type UserInfoDTO struct {
	ID        uint      `json:"ID"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
}
