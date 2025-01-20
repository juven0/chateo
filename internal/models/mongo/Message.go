package mongomodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ConversationID primitive.ObjectID `bson:"conversation_id" json:"conversationId"`
	SenderID       primitive.ObjectID `bson:"sender_id" json:"senderId"`
	//temporary
	To 			   primitive.ObjectID `bson:"recever_id" json:"receverId"`
	Content        string             `bson:"content" json:"content"`
	Type           string             `bson:"type" json:"type" validate:"oneof=text image file"`
	ReadBy         []ReadReceipt      `bson:"read_by" json:"readBy"`
	Status         string             `bson:"status" json:"status" validate:"oneof=sent delivered read"`
	CreatedAt      time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updatedAt"`
}