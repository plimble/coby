package coby

//go:generate mockery -name=Store -inpkg

type Store interface {
	Create(token *Token) error
	Delete(token string) error
	Get(token string) (*Token, error)
}
