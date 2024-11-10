package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	ID            primitive.ObjectID `bson:"_id"`
	Token         *string             `json:"token"`
	Refresh_token *string             `json:"refresh_token"`
	Reset_token *string             `json:"reset_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	Expiry    time.Time          `json:"expiry"`
	TTL    time.Time          `json:"ttl"`
	User_id       string             `json:"user_id"`
	Scope *string `json:"scope"`
	// TODO: add more fields
}