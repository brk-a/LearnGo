package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string             `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string             `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string             `json:"Password" validate:"required,min=12"`
	Email         *string             `json:"email" validate:"required,email"`
	Avatar        *string             `json:"avatar"`
	Phone         *string             `json:"phone" validate:"required,min=8"`
	Token         *string             `json:"token"`
	Refresh_token *string             `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
	// Last_login time.Time `json:"last_login"`
	// Role       string             `json:"role" validate:"required,eq=admin|eq=manager|eq=waiter|eq=kitchen"`
}
