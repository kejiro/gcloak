package gcloak

import "errors"

var (
	ErrAccessDenied = errors.New("access denied")
	ErrNotFound     = errors.New("not found")
)
