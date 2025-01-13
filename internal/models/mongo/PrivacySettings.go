package mongomodels

type PrivacySettings struct {
	ShowOnlineStatus bool `bson:"show_online_status" json:"showOnlineStatus"`
	ShowLastSeen     bool `bson:"show_last_seen" json:"showLastSeen"`
}