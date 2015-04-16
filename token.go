package coby

//go:generate msgp

type Token struct {
	Token  string   `json:"i" bson:"_id" redis:"i" msg:"t"`
	Data   []string `json:"d,omitempty" bson:"d,omitempty" redis:"d" msg:"d"`
	Expire int64    `json:"e" bson:"e" redis:"e" msg:"e"`
}
