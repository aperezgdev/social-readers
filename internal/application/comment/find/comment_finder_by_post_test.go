package finder_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	finder "github.com/aperezgdev/social-readers-api/internal/application/comment/find"
	"github.com/aperezgdev/social-readers-api/internal/domain/errors"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
)

type testSuite struct {
	mockCommentRepository *repository.MockCommentRepository
	mockPostRepository    *repository.MockPostRepository
	comentFinderByPost    *finder.CommentFinderByPost
}

func setupTest() *testSuite {
	mockCommentRepository := repository.MockCommentRepository{}
	mockPostRepository := repository.MockPostRepository{}
	commentFinderByPost := finder.NewCommentFinderByPost(
		slog.Default(),
		&mockCommentRepository,
		&mockPostRepository,
	)

	return &testSuite{
		mockPostRepository:    &mockPostRepository,
		mockCommentRepository: &mockCommentRepository,
		comentFinderByPost:    &commentFinderByPost,
	}
}

func TestCommentFinder(t *testing.T) {
	t.Parallel()

	t.Run("should find comments", func(t *testing.T) {
		suite := setupTest()
		suite.mockPostRepository.On("Find", mock.Anything, mock.Anything).
			Return(models.Post{}, nil).
			Once()
		suite.mockCommentRepository.On("FindByPost", mock.Anything, mock.Anything).
			Return(make([]models.Comment, 10), nil).Once()

		comments, err := suite.comentFinderByPost.Run(context.Background(), "1")
		assert.Nil(t, err)
		assert.True(t, len(comments) == 10)

		suite.mockCommentRepository.AssertExpectations(t)
		suite.mockPostRepository.AssertExpectations(t)
	})

	t.Run("should fail no exists post", func(t *testing.T) {
		suite := setupTest()
		suite.mockPostRepository.On("Find", mock.Anything, mock.Anything).
			Return(models.Post{}, errors.ErrNotExistPost).
			Once()

		_, err := suite.comentFinderByPost.Run(context.Background(), "1")
		assert.ErrorIs(t, err, errors.ErrNotExistPost)

		suite.mockPostRepository.AssertExpectations(t)
	})
}
