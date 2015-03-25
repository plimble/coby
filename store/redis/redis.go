package redis

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/plimble/coby"
	"time"
)

type Store struct {
	*redis.Pool
}

func NewStore(r *redis.Pool) *Store {
	return &Store{r}
}

func NewPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (s *Store) getC() redis.Conn {
	return s.Pool.Get()
}

func (s *Store) Create(token *coby.Token) error {
	b, err := json.Marshal(token)
	if err != nil {
		return err
	}

	if _, err := s.getC().Do("SET", token.Token, string(b)); err != nil {
		return err
	}
	return nil
}

func (s *Store) Delete(id string) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	if _, err = s.getC().Do("SET", id, string(b)); err != nil {
		return err
	}
	return nil
}

func (s *Store) Get(tokenID string) (*coby.Token, error) {
	r, err := s.getC().Do("GET", tokenID)
	if err != nil {
		return nil, err
	}

	var t *coby.Token
	str, err := redis.Bytes(r, nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(str, &t)

	return t, err
}
