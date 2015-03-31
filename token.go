package coby

type Token struct {
	Token  string   `json:"i" bson:"_id" redis:"i"`
	Data   []string `json:"d,omitempty" bson:"d,omitempty" redis:"d"`
	Expire int64    `json:"e" bson:"e" redis:"e"`
}
