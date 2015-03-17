package coby

type benchStore struct{}

func (s *benchStore) Create(id string, v interface{}) error {
	return nil
}

func (s *benchStore) Update(id string, v interface{}) error {
	return nil
}

func (s *benchStore) Get(tokenID string) (*Token, error) {
	return &Token{}, nil
}
