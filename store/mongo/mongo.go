package mongo

import (
	"github.com/plimble/coby"
	"gopkg.in/mgo.v2"
)

type Store struct {
	db      string
	c       string
	session *mgo.Session
}

func NewStore(session *mgo.Session) *Store {
	return &Store{
		db:      "coby",
		c:       "token",
		session: session,
	}
}

func (s *Store) getC(session *mgo.Session) *mgo.Collection {
	return session.DB(s.db).C(s.c)
}

func (s *Store) Create(token *coby.Token) error {
	session := s.session.Clone()
	defer session.Close()

	return s.getC(session).Insert(token)
}

func (s *Store) Delete(id string) error {
	session := s.session.Clone()
	defer session.Close()

	return s.getC(session).RemoveId(id)
}

func (s *Store) Get(tokenID string) (*coby.Token, error) {
	session := s.session.Clone()
	defer session.Close()

	var token *coby.Token
	err := s.getC(session).FindId(tokenID).One(&token)

	return token, err
}
