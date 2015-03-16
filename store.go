package coby

import (
	"github.com/plimble/crud"
)

//go:generate mockery -file=store_mock.go -name=Store -inpkg

type Store interface {
	crud.CRUD
	Get(tokenID string) (*Token, error)
}
