package mongomodels

type UserPreferences struct {
	Notifications      bool            `bson:"notifications" json:"notifications"`
	EmailNotifications bool            `bson:"email_notifications" json:"emailNotifications"`
	Theme              string          `bson:"theme" json:"theme" validate:"oneof=light dark"`
	Privacy            PrivacySettings `bson:"privacy" json:"privacy"`
}