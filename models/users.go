package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User: data of user
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `bson:"created_At" json:"created_at"`
	UpdatedAt time.Time          `bson:"created_at" json:"updated_at,omitempty"`
}

// Users: List of users
type Users []*User
