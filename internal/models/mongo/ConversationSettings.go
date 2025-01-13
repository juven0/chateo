package mongomodels

type ConversationSettings struct {
	Notifications bool `bson:"notifications" json:"notifications"`
	Encrypted     bool `bson:"encrypted" json:"encrypted"`
}