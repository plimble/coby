package coby

import (
	"encoding/json"
	"github.com/plimble/moment"
	"github.com/plimble/unik"
)

//go:generate mockery -file=coby_mock.go -name=Service -inpkg=false

type Service interface {
	CreateToken(v interface{}) (*Token, error)
	GetToken(tokenID string) (*Token, error)
	UseToken(tokenID string) error
}

type CobyService struct {
	Store  Store
	Unik   unik.Generator
	Moment moment.Time
}

func NewService(store Store, unik unik.Generator, moment moment.Time) *CobyService {
	return &CobyService{
		Store:  store,
		Unik:   unik,
		Moment: moment,
	}
}

func (c *CobyService) CreateToken(v interface{}) (*Token, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	t := &Token{
		ID:     c.Unik.Generate(),
		Data:   string(b),
		Expire: c.Moment.Now(),
		Used:   false,
	}

	err = c.Store.Create(t.ID, t)
	return t, err
}

func (c *CobyService) GetToken(tokenID string) (*Token, error) {
	return c.Store.Get(tokenID)
}

func (c *CobyService) UseToken(tokenID string) error {
	var token *Token
	token, err := c.Store.Get(tokenID)
	if err != nil {
		return err
	}

	if token.Used == true {
		return nil
	}

	token.Used = true
	return c.Store.Update(tokenID, token)
}
