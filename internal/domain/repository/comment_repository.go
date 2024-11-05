package repository

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	post_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/post"
)

type CommentRepository interface {
	FindByPost(context.Context, post_vo.PostId) ([]models.Comment, error)
	Save(context.Context, models.Comment) error
}

type MockCommentRepository struct {
	mock.Mock
}

func NewMockCommentRepository() MockCommentRepository {
	return MockCommentRepository{}
}

func (m *MockCommentRepository) FindByPost(
	ctx context.Context,
	postId post_vo.PostId,
) ([]models.Comment, error) {
	args := m.Called(ctx, postId)

	return args.Get(0).([]models.Comment), args.Error(1)
}

func (m *MockCommentRepository) Save(ctx context.Context, comment models.Comment) error {
	args := m.Called(ctx, comment)

	return args.Error(0)
}
