package coby

import "github.com/stretchr/testify/mock"

type MockStore struct {
	mock.Mock
}

func NewMockStore() *MockStore {
	return &MockStore{}
}

func (m *MockStore) Insert(v interface{}) error {
	ret := m.Called(v)

	r0 := ret.Error(0)

	return r0
}
func (m *MockStore) Delete(id interface{}) error {
	ret := m.Called(id)

	r0 := ret.Error(0)

	return r0
}
func (m *MockStore) Update(id interface{}, v map[string]interface{}) error {
	ret := m.Called(id, v)

	r0 := ret.Error(0)

	return r0
}
func (m *MockStore) UpdateAll(id interface{}, v interface{}) error {
	ret := m.Called(id, v)

	r0 := ret.Error(0)

	return r0
}
func (m *MockStore) Upsert(id interface{}, v interface{}) error {
	ret := m.Called(id, v)

	r0 := ret.Error(0)

	return r0
}
func (m *MockStore) Exist(id interface{}) (bool, error) {
	ret := m.Called(id)

	r0 := ret.Get(0).(bool)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *MockStore) Get(tokenID string) (*Token, error) {
	ret := m.Called(tokenID)

	r0 := ret.Get(0).(*Token)
	r1 := ret.Error(1)

	return r0, r1
}
