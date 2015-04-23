package aerospike

import (
	"github.com/aerospike/aerospike-client-go"
	"github.com/plimble/aerosingle"
	"github.com/plimble/coby"
)

type Store struct {
	as *aerosingle.Client
}

func NewStore(asClient *aerosingle.Client) *Store {
	return &Store{asClient}
}

func (s *Store) Create(token *coby.Token) error {
	policy := aerospike.NewWritePolicy(0, 0)
	policy.RecordExistsAction = aerospike.CREATE_ONLY

	return s.as.PutMsgPack(policy, "coby", token.Token, token)
}

func (s *Store) Delete(id string) error {
	return s.as.Delete(nil, "coby", id)
}

func (s *Store) Get(tokenID string) (*coby.Token, error) {
	var t coby.Token

	err := s.as.GetMsgPack(nil, "coby", tokenID, &t)

	return &t, err
}
