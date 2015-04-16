package aerospike

import (
	"github.com/aerospike/aerospike-client-go"
	"github.com/plimble/coby"
	"github.com/plimble/utils/assingle"
)

type Store struct {
	as *assingle.ASSingle
}

func NewStore(asClient *aerospike.Client, ns string) *Store {
	return &Store{assingle.New(asClient, ns)}
}

func (s *Store) Create(token *coby.Token) error {
	policy := aerospike.NewWritePolicy(0, 0)
	policy.RecordExistsAction = aerospike.CREATE_ONLY

	return s.as.Put(policy, "coby", token.Token, token)
}

func (s *Store) Delete(id string) error {
	return s.as.Delete(nil, "coby", id)
}

func (s *Store) Get(tokenID string) (*coby.Token, error) {
	var t coby.Token

	err := s.as.Get(nil, "coby", tokenID, &t)

	return &t, err
}
