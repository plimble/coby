package coby

import "github.com/stretchr/testify/mock"

type MockStore struct {
	mock.Mock
}

func NewMockStore() *MockStore {
	return &MockStore{}
}

func (m *MockStore) Create(id string, v interface{}) error {
	ret := m.Called(id, v)

	r0 := ret.Error(0)

	return r0
}
func (m *MockStore) Update(id string, v interface{}) error {
	ret := m.Called(id, v)

	r0 := ret.Error(0)

	return r0
}
func (m *MockStore) Get(tokenID string) (*Token, error) {
	ret := m.Called(tokenID)

	r0 := ret.Get(0).(*Token)
	r1 := ret.Error(1)

	return r0, r1
}
