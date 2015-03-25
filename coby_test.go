package coby

import (
	"errors"
	"github.com/plimble/moment/mock_moment"
	"github.com/plimble/unik/mock_unik"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockCoby struct {
	store  *MockStore
	unik   *mock_unik.MockGenerator
	moment *mock_moment.MockTime
}

func setupCoby() (*CobyService, *mockCoby) {
	store := NewMockStore()
	unik := mock_unik.NewMockGenerator()
	moment := mock_moment.NewMockTime()
	c := NewService(store, time.Second*time.Duration(30))
	mock := &mockCoby{store, unik, moment}
	c.unik = unik
	c.moment = moment

	return c, mock
}

func TestCreate(t *testing.T) {
	s, m := setupCoby()

	expToken := &Token{
		Token:  "123",
		Expire: int64(100),
		Data: map[string]interface{}{
			"email": "test@test.com",
		},
	}

	m.unik.On("Generate").Return("123")
	m.moment.On("AddNowUnix", s.expires).Return(int64(100))
	m.store.On("Create", expToken).Return(nil)

	token, err := s.Create(expToken.Data)
	m.unik.AssertExpectations(t)
	m.moment.AssertExpectations(t)
	m.store.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expToken, token)
}

func TestVerify(t *testing.T) {
	s, m := setupCoby()

	expToken := &Token{
		Token:  "123",
		Expire: time.Now().UTC().Add(time.Second * 20).Unix(),
		Data: map[string]interface{}{
			"email": "test@test.com",
		},
	}

	m.store.On("Get", expToken.Token).Return(expToken, nil)
	m.moment.On("IsExpireUnix", expToken.Expire).Return(false)

	token, err := s.Verify(expToken.Token)
	m.moment.AssertExpectations(t)
	m.store.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expToken, token)
}

func TestVerify_Expired(t *testing.T) {
	s, m := setupCoby()

	expToken := &Token{
		Token:  "123",
		Expire: time.Now().UTC().Add(time.Second * -20).Unix(),
		Data: map[string]interface{}{
			"email": "test@test.com",
		},
	}

	m.store.On("Get", expToken.Token).Return(expToken, nil)
	m.moment.On("IsExpireUnix", expToken.Expire).Return(true)

	token, err := s.Verify(expToken.Token)
	m.store.AssertExpectations(t)
	assert.Equal(t, err, errTokenExpired)
	assert.Nil(t, token)
}

func TestVerify_NotFound(t *testing.T) {
	s, m := setupCoby()

	expToken := &Token{
		Token:  "123",
		Expire: time.Now().UTC().Add(time.Second * -20).Unix(),
		Data: map[string]interface{}{
			"email": "test@test.com",
		},
	}

	m.store.On("Get", expToken.Token).Return(nil, errors.New("error"))

	token, err := s.Verify(expToken.Token)
	m.store.AssertExpectations(t)
	assert.Equal(t, err, errInvalidToken)
	assert.Nil(t, token)
}

func TestDelete(t *testing.T) {
	s, m := setupCoby()

	m.store.On("Delete", "123").Return(nil)

	err := s.Delete("123")
	assert.NoError(t, err)
}

func BenchmarkCreate(b *testing.B) {
	store := &benchStore{}
	c := NewService(store, time.Second*30)
	b.ReportAllocs()
	b.ResetTimer()

	data := map[string]interface{}{
		"name": "123",
		"pass": "321",
	}

	for n := 0; n < b.N; n++ {
		c.Create(data)
	}
}

func BenchmarkVerify(b *testing.B) {
	store := &benchStore{}
	c := NewService(store, time.Second*30)
	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		c.Verify("123")
	}
}

func BenchmarkDelete(b *testing.B) {
	store := &benchStore{}
	c := NewService(store, time.Second*30)
	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		c.Delete("123")
	}
}
