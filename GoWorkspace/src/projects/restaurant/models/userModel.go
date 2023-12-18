package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string            `json:"Password" validate:"required,min=6"`
	Email         *string            `json:"Email" validate:"required"`
	Avatar        *string            `json:"Avatar"`
	Phone         *string            `json:"Phone" validate:"required"`
	Token         *string            `json:"Token"`
	Refresh_token *string            `json:"Refresh_token"`
	Created_at    time.Time          `json:"Created_at"`
	Updated_at    time.Time          `json:"Updated_at"`
	User_id       string             `json:"user_id"`
}
