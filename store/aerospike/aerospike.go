package aerospike

import (
	"github.com/plimble/aero"
	"github.com/plimble/coby"
)

type Store struct {
	as  *aero.Client
	ns  string
	set string
}

func NewStore(asClient *aero.Client, ns string) *Store {
	return &Store{asClient, ns, "coby"}
}

func (s *Store) Create(token *coby.Token) error {
	policy := aero.NewWritePolicy(0, 0)
	policy.RecordExistsAction = aero.CREATE_ONLY

	binToken := aero.NewBin("token", token.Token)
	binData := aero.NewBin("data", token.Data)
	binExpire := aero.NewBin("expire", token.Expire)

	return s.as.PutBins(policy, s.ns, s.set, token.Token, binToken, binData, binExpire)
}

func (s *Store) Delete(id string) error {
	return s.as.Delete(nil, s.ns, s.set, id)
}

func (s *Store) Get(tokenID string) (*coby.Token, error) {
	rec, err := s.as.Get(nil, s.ns, s.set, tokenID, "token", "data", "expire")
	if err != nil {
		return nil, err
	}

	c := &coby.Token{}
	c.Token = rec.Bins["token"].(string)

	data := rec.Bins["data"].([]interface{})
	c.Data = make([]string, len(data))
	for i := 0; i < len(data); i++ {
		c.Data[i] = data[i].(string)
	}

	c.Expire = int64(rec.Bins["expire"].(int))

	return c, nil
}
