package gouxp

import "github.com/pkg/errors"

var (
	ErrDifferentAddr error = errors.New("different remote addr.")
)