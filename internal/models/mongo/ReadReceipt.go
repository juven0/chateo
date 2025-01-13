package mongomodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadReceipt struct {
	UserID primitive.ObjectID `bson:"user_id" json:"userId"`
	ReadAt time.Time          `bson:"read_at" json:"readAt"`
}