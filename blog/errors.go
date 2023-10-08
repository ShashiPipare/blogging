package blog

import (
	"errors"
)

var (
	ErrInvalidUserID   = errors.New("invalid user_id")
	ErrInvalidBlogID   = errors.New("invalid blog_id")
	ErrInvalidString   = errors.New("invalid string")
	ErrInvalidNotFound = errors.New("invalid key/key not found")
	ErrEmptyArray      = errors.New("empty array not allowed")
)
