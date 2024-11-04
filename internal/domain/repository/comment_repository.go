package repository

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	comment_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/comment"
)

type CommentRepository interface {
	Find(context.Context, comment_vo.ComentId) (models.Comment, error)
	Save(context.Context, models.Comment) error
}

type MockCommentRepository struct {
	mock.Mock
}

func NewMockCommentRepository() MockCommentRepository {
	return MockCommentRepository{}
}

func (m *MockCommentRepository) Find(
	ctx context.Context,
	id comment_vo.ComentId,
) (models.Comment, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(models.Comment), args.Error(1)
}

func (m *MockCommentRepository) Save(ctx context.Context, comment models.Comment) error {
	args := m.Called(ctx, comment)

	return args.Error(0)
}
