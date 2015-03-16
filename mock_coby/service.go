package mock_coby

import "bitbucket.org/plimble/coby"
import "github.com/stretchr/testify/mock"

type MockService struct {
	mock.Mock
}

func NewMockService() *MockService {
	return &MockService{}
}

func (m *MockService) CreateToken(v interface{}) (*coby.Token, error) {
	ret := m.Called(v)

	r0 := ret.Get(0).(*coby.Token)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *MockService) GetToken(tokenID string) (*coby.Token, error) {
	ret := m.Called(tokenID)

	r0 := ret.Get(0).(*coby.Token)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *MockService) UseToken(tokenID string) error {
	ret := m.Called(tokenID)

	r0 := ret.Error(0)

	return r0
}
