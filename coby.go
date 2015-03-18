package coby

import (
	"encoding/json"
	"github.com/plimble/moment"
	"github.com/plimble/unik"
)

//go:generate mockery -name=Service -inpkg=false

type Service interface {
	CreateToken(v interface{}) (*Token, error)
	GetToken(tokenID string) (*Token, error)
	UseToken(tokenID string) error
}

type CobyService struct {
	store  Store
	unik   unik.Generator
	moment moment.Time
}

func NewService(store Store) *CobyService {
	return &CobyService{
		store:  store,
		unik:   unik.NewSnowflake(1),
		moment: moment.New(),
	}
}

func (c *CobyService) CreateToken(v interface{}) (*Token, error) {
	t := &Token{
		ID:     c.unik.Generate(),
		Expire: c.moment.Now(),
		Used:   false,
	}

	if v != nil {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		t.Data = string(b)
	}

	err := c.store.Create(t.ID, t)
	return t, err
}

func (c *CobyService) GetToken(tokenID string) (*Token, error) {
	return c.store.Get(tokenID)
}

func (c *CobyService) UseToken(tokenID string) error {
	var token *Token
	token, err := c.store.Get(tokenID)
	if err != nil {
		return err
	}

	if token.Used == true {
		return nil
	}

	token.Used = true
	return c.store.Update(tokenID, token)
}
