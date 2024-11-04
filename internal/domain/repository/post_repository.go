package repository

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	post_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/post"
)

type PostRepository interface {
	Find(context.Context, post_vo.PostId) (models.Post, error)
	Save(context.Context, models.Post) error
}

type MockPostRepository struct {
	mock.Mock
}

func (m *MockPostRepository) Find(ctx context.Context, id post_vo.PostId) (models.Post, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(models.Post), args.Error(1)
}

func (m *MockPostRepository) Save(ctx context.Context, post models.Post) error {
	args := m.Called(ctx, post)

	return args.Error(0)
}
