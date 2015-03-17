package mongo

import (
	"bitbucket.org/plimble/coby"
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

func (s *Store) Create(id string, v interface{}) error {
	session := s.session.Clone()
	defer session.Close()

	return s.getC(session).Insert(v)
}

func (s *Store) Update(id string, v interface{}) error {
	session := s.session.Clone()
	defer session.Close()

	return s.getC(session).UpdateId(id, v)
}

func (s *Store) Get(tokenID string) (*coby.Token, error) {
	session := s.session.Clone()
	defer session.Close()

	var token *coby.Token
	err := s.getC(session).FindId(tokenID).One(&token)

	return token, err
}
