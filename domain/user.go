package domain

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInternalServer     = errors.New("internal server error")
)

type User struct {
	ID       string
	Username string
	Password string
}

type UserRepository interface {
	FindByUsername(username string) (*User, error)
}
