package mongomodels

type User struct {
	Id        string `json:"id" bson:"_id"`
	UserName  string `json:"username" bson:"username"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	Status    bool   `json:"status" bson:"status"`
	LastSeen  string `json:"lastseen" bson:"lastseen"`
	CreatedAt string `json:"createdat" bson:"createdat"`
	UpdateAt  string `json:"updateat" bson:"updateat"`
}