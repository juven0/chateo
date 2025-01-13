package mongomodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRelation struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"userId"`
	ContactID primitive.ObjectID `bson:"contact_id" json:"contactId"`
	Status    string             `bson:"status" json:"status" validate:"oneof=pending accepted blocked"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}