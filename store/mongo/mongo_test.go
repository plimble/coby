package mongo

import (
	"bitbucket.org/plimble/coby"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"os"
	"testing"
	"time"
)

func init() {
	time.Local = time.UTC
}

func setup() *Store {
	session, err := mgo.Dial(os.Getenv("MONGO"))
	if err != nil {
		panic(err)
	}

	return NewStore(session)
}

func tearDown(s *Store) {
	s.session.DB(s.db).DropDatabase()
	s.session.Close()
}

func generateToken(id string) *coby.Token {
	return &coby.Token{
		ID:     id,
		Data:   "",
		Expire: time.Now().Truncate(time.Second),
		Used:   false,
	}
}

type testData struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func TestMongoStore(t *testing.T) {
	s := setup()

	d := &testData{
		Name:  "Tickets",
		Owner: "Admin",
	}
	j, err := json.Marshal(d)
	assert.NoError(t, err)

	token := generateToken("1")
	token.Data = string(j)

	err = s.Create(token.ID, token)
	assert.NoError(t, err)

	result, err := s.Get(token.ID)
	assert.NoError(t, err)
	assert.Equal(t, token, result)

	token.Used = true
	err = s.Update(token.ID, token)
	assert.NoError(t, err)

	result, err = s.Get(token.ID)
	assert.NoError(t, err)
	assert.Equal(t, token, result)

	tearDown(s)
}
