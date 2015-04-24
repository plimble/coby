package aerospike

import (
	"github.com/plimble/aero"
	"github.com/plimble/coby"
	"github.com/plimble/rand"
	"github.com/stretchr/testify/suite"
	"testing"
)

type StoreSuite struct {
	s *Store
	suite.Suite
}

func TestStoreSuite(t *testing.T) {
	suite.Run(t, &StoreSuite{})
}

func genToken() *coby.Token {
	return &coby.Token{
		Token:  rand.Digits(4),
		Data:   []string{rand.City(), rand.City(), rand.City()},
		Expire: int64(rand.Number(10, 1000)),
	}
}

func (t *StoreSuite) SetupSuite() {
	t.s = NewStore(aero.NewClient("192.168.99.112", 3000), "test")
}

func (t *StoreSuite) TestCrud() {
	expt := genToken()

	err := t.s.Create(expt)
	t.NoError(err)

	token, err := t.s.Get(expt.Token)
	t.NoError(err)
	t.Equal(expt, token)

	err = t.s.Delete(expt.Token)
	t.NoError(err)

	token, err = t.s.Get(expt.Token)
	t.Nil(token)
	t.Equal(aero.ErrNotFound, err)
}
