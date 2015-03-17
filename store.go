package coby

//go:generate mockery -file=store_mock.go -name=Store -inpkg

type Store interface {
	Create(id string, v interface{}) error
	Update(id string, v interface{}) error
	Get(tokenID string) (*Token, error)
}
