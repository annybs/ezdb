package ezdb

import "errors"

// High-level EZ DB error.
// These are not exhaustive and your chosen implementation of Collection may produce its own errors.
var (
	ErrClosed     = errors.New("collection is closed")
	ErrInvalidKey = errors.New("invalid key")
	ErrNotFound   = errors.New("not found")
	ErrReleased   = errors.New("iterator has been released")
)
