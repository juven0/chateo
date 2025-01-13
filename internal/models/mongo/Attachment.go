package mongomodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Attachment struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	MessageID      primitive.ObjectID `bson:"message_id" json:"messageId"`
	ConversationID primitive.ObjectID `bson:"conversation_id" json:"conversationId"`
	Type           string             `bson:"type" json:"type" validate:"oneof=image video document"`
	URL            string             `bson:"url" json:"url"`
	Size           int64              `bson:"size" json:"size"`
	CreatedAt      time.Time          `bson:"created_at" json:"createdAt"`
}