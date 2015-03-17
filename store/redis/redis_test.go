package redis

import (
	"bitbucket.org/plimble/coby"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func generateToken(id string) *coby.Token {
	return &coby.Token{
		ID:     id,
		Data:   "",
		Expire: time.Now().Truncate(time.Second),
		Used:   false,
	}
}

func TestRedisStore(t *testing.T) {
	p := NewPool(os.Getenv("REDIS"))
	s := NewStore(p)

	token := generateToken("1")
	err := s.Create(token.ID, token)
	assert.NoError(t, err)

	result, err := s.Get("1")
	assert.NoError(t, err)
	assert.Equal(t, token, result)

	token.Used = true
	err = s.Update(token.ID, token)
	assert.NoError(t, err)

	result, err = s.Get("1")
	assert.NoError(t, err)
	assert.True(t, token.Used)

	s.getC().Flush()
}
