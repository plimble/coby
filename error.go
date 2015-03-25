package coby

import (
	"github.com/plimble/utils/errors2"
)

var (
	errInvalidToken = errors2.NewNotFound("Invalid Token")
	errTokenExpired = errors2.NewNotFound("Token is expired")
)
