package coby

import (
	"encoding/json"
	"time"
)

type Token struct {
	ID     string    `json:"id" bson:"_id"`
	Data   string    `json:"data" bson:"d"`
	Expire time.Time `json:"expire" bson:"e"`
	Used   bool      `json:"used" bson:"u"`
}

func (t *Token) IsExpire() bool {
	if t.Expire.Before(time.Now().UTC()) {
		return true
	}
	return false
}

func (t *Token) GetData(v interface{}) error {
	return json.Unmarshal([]byte(t.Data), &v)
}
