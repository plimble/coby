package coby

import (
	"code.google.com/p/gomock/gomock"
	"encoding/json"
	"errors"
	"github.com/plimble/moment/mock_moment"
	"github.com/plimble/unik/mock_unik"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type fake struct {
	ctrl   *gomock.Controller
	store  *MockStore
	unik   *mock_unik.MockGenerator
	moment *mock_moment.MockTime
}

func newFake(t *testing.T) *fake {
	ctrl := gomock.NewController(t)
	store := NewMockStore()
	unik := mock_unik.NewMockGenerator(ctrl)
	moment := mock_moment.NewMockTime(ctrl)
	return &fake{
		ctrl:   ctrl,
		store:  store,
		unik:   unik,
		moment: moment,
	}
}

func generateToken(id string) *Token {
	return &Token{
		ID:     id,
		Data:   "",
		Expire: time.Now(),
		Used:   false,
	}
}

type testData struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func TestCreateToken(t *testing.T) {
	f := newFake(t)
	c := NewService(f.store, f.unik, f.moment)

	d := &testData{
		Name:  "Tickets",
		Owner: "Admin",
	}

	token := generateToken("1")
	jsonStr, err := json.Marshal(&d)
	assert.NoError(t, err)

	token.Data = string(jsonStr)

	f.unik.EXPECT().Generate().Return(token.ID)
	f.moment.EXPECT().Now().Return(token.Expire)
	f.store.On("Create", token.ID, token).Return(nil)

	result, err := c.CreateToken(d)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	m, err := json.Marshal(d)
	result.Data = string(m)
	v := &testData{}
	err = result.GetData(v)

	assert.NoError(t, err)
	assert.Equal(t, d, v)
}

func TestCreateWithNil(t *testing.T) {
	f := newFake(t)
	c := NewService(f.store, f.unik, f.moment)

	token := generateToken("1")
	token.Data = "null"

	f.unik.EXPECT().Generate().Return(token.ID)
	f.moment.EXPECT().Now().Return(token.Expire)
	f.store.On("Create", token.ID, token).Return(nil)

	_, err := c.CreateToken(nil)
	assert.NoError(t, err)
}

func TestGetToken(t *testing.T) {
	f := newFake(t)
	c := NewService(f.store, f.unik, f.moment)

	token := generateToken("1")
	f.store.On("Get", "1").Return(token, nil)

	result, err := c.GetToken("1")
	assert.NoError(t, err)
	assert.Equal(t, result, token)
}

func TestUseToken(t *testing.T) {
	f := newFake(t)
	c := NewService(f.store, f.unik, f.moment)

	token := generateToken("1")
	f.store.On("Get", "1").Return(token, nil)
	token.Used = true
	f.store.On("Update", "1", token).Return(nil)

	err := c.UseToken("1")
	assert.NoError(t, err)
}

func TestUseTokenAlready(t *testing.T) {
	f := newFake(t)
	c := NewService(f.store, f.unik, f.moment)

	token := generateToken("1")
	token.Used = true
	f.store.On("Get", "1").Return(token, nil)

	err := c.UseToken("1")
	assert.NoError(t, err)
}

func TestUseTokenNotFound(t *testing.T) {
	f := newFake(t)
	c := NewService(f.store, f.unik, f.moment)

	f.store.On("Get", "1").Return(&Token{}, errors.New("not found"))

	err := c.UseToken("1")
	assert.Error(t, err)
}

func TestTokenIsExpire(t *testing.T) {
	token := generateToken("1")
	b := token.IsExpire()
	assert.True(t, b)

	token.Expire = time.Now().AddDate(1, 0, 0)
	b = token.IsExpire()
	assert.False(t, b)

}
