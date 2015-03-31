package coby

import (
	"github.com/plimble/moment"
	"github.com/plimble/unik"
	"time"
)

//go:generate mockery -name=Service

type Service interface {
	Create(v interface{}) (*Token, error)
	Verify(token string) (*Token, error)
	Delete(tokenID string) error
}

type CobyService struct {
	store   Store
	unik    unik.Generator
	moment  moment.Time
	expires time.Duration
}

func NewService(store Store, expires time.Duration) *CobyService {
	return &CobyService{
		store:   store,
		unik:    unik.NewBSON(),
		moment:  moment.New(),
		expires: expires,
	}
}

func (c *CobyService) Create(v interface{}) (*Token, error) {
	t := &Token{
		Token:  c.unik.Generate(),
		Expire: c.moment.AddNowUnix(c.expires),
		Data:   v,
	}

	err := c.store.Create(t)
	return t, err
}

func (c *CobyService) Verify(token string) (*Token, error) {
	t, err := c.store.Get(token)
	if err != nil {
		return nil, errInvalidToken
	}

	if c.moment.IsExpireUnix(t.Expire) {
		return nil, errTokenExpired
	}

	return t, nil
}

func (c *CobyService) Delete(token string) error {
	return c.store.Delete(token)
}
