package repository

import (
	"github.com/sleepiie/login-system/domain"

	"golang.org/x/crypto/bcrypt"
)

type mockUserRepository struct {
	users map[string]*domain.User
}

func NewMockUserRepository() domain.UserRepository {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	return &mockUserRepository{
		users: map[string]*domain.User{
			"admin": {
				ID:       "1",
				Username: "admin",
				Password: string(hashedPassword),
			},
		},
	}
}

func (m *mockUserRepository) FindByUsername(username string) (*domain.User, error) {
	user, exists := m.users[username]
	if !exists {
		return nil, domain.ErrUserNotFound
	}
	return user, nil
}
