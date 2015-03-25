package coby

import (
	"time"
)

type benchStore struct{}

func (s *benchStore) Create(token *Token) error {
	return nil
}

func (s *benchStore) Delete(id string) error {
	return nil
}

func (s *benchStore) Get(tokenID string) (*Token, error) {
	return &Token{
		Expire: time.Now().UTC().Add(time.Second * 100).Unix(),
	}, nil
}
