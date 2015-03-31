package mock_coby

import "github.com/plimble/coby"
import "github.com/stretchr/testify/mock"

type MockService struct {
	mock.Mock
}

func NewMockService() *MockService {
	return &MockService{}
}

func (m *MockService) Create(v map[string]interface{}) (*coby.Token, error) {
	ret := m.Called(v)

	r0 := ret.Get(0).(*coby.Token)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *MockService) Verify(token string) (*coby.Token, error) {
	ret := m.Called(token)

	r0 := ret.Get(0).(*coby.Token)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *MockService) Delete(tokenID string) error {
	ret := m.Called(tokenID)

	r0 := ret.Error(0)

	return r0
}
