package coby

import "github.com/stretchr/testify/mock"

type MockStore struct {
	mock.Mock
}

func NewMockStore() *MockStore {
	return &MockStore{}
}

func (m *MockStore) Create(token *Token) error {
	ret := m.Called(token)

	r0 := ret.Error(0)

	return r0
}
func (m *MockStore) Delete(token string) error {
	ret := m.Called(token)

	r0 := ret.Error(0)

	return r0
}
func (m *MockStore) Get(token string) (*Token, error) {
	ret := m.Called(token)

	var r0 *Token
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*Token)
	}
	r1 := ret.Error(1)

	return r0, r1
}
