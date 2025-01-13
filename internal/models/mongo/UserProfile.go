package mongomodels

type UserProfile struct {
	FirstName string `bson:"first_name" json:"firstName" validate:"max=50"`
	LastName  string `bson:"last_name" json:"lastName" validate:"max=50"`
	Avatar    string `bson:"avatar" json:"avatar"`
	Bio       string `bson:"bio" json:"bio" validate:"max=500"`
	Language  string `bson:"language" json:"language" validate:"len=2"`
	Timezone  string `bson:"timezone" json:"timezone"`
}