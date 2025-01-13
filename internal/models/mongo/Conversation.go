package mongomodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Conversation struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Type         string               `bson:"type" json:"type" validate:"oneof=private group"`
	Name         string               `bson:"name,omitempty" json:"name"`
	Participants []primitive.ObjectID `bson:"participants" json:"participants"`
	CreatedBy    primitive.ObjectID   `bson:"created_by" json:"createdBy"`
	CreatedAt    time.Time            `bson:"created_at" json:"createdAt"`
	LastMessage  time.Time            `bson:"last_message" json:"lastMessage"`
	Settings     ConversationSettings `bson:"settings" json:"settings"`
}