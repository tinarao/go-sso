package models

import "gorm.io/gorm"

var Roles [3]string = [3]string{"user", "admin", "moderator"}

type User struct {
	gorm.Model
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}
