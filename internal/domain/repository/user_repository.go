package repository

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
)

type UserRepository interface {
	Find(context.Context, user_vo.UserId) (models.User, error)
	Save(context.Context, models.User) error
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Find(ctx context.Context, id user_vo.UserId) (models.User, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) Save(ctx context.Context, user models.User) error {
	args := m.Called(ctx, user)

	return args.Error(0)
}
