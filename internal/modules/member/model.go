package member

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Member struct {
		Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		FullName  string             `json:"fullName"`
		Username  string             `json:"username" bson:"username"`
		Password  string             `json:"password" bson:"password"`
		Email     string             `json:"email" bson:"email"`
		CreatedAt time.Time          `json:"createdAt" bson:"created_at"`
		UpdatedAt time.Time          `json:"updatedAt" bson:"updated_at"`
	}
)
