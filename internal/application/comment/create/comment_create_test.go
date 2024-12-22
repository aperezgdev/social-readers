package create_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/social-readers-api/internal/application/comment/create"
	"github.com/aperezgdev/social-readers-api/internal/domain/errors"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
)

type testSuite struct {
	commentRepository *repository.MockCommentRepository
	userRepository    *repository.MockUserRepository
	postRepository    *repository.MockPostRepository
	commentCreator    *create.CommentCreator
}

func setupTest() *testSuite {
	mockCommentRepo := &repository.MockCommentRepository{}
	mockUserRepo := &repository.MockUserRepository{}
	mockPostRepo := &repository.MockPostRepository{}

	creator := create.NewCommentCreator(
		slog.Default(),
		mockCommentRepo,
		mockUserRepo,
		mockPostRepo,
	)

	return &testSuite{
		commentRepository: mockCommentRepo,
		userRepository:    mockUserRepo,
		postRepository:    mockPostRepo,
		commentCreator:    &creator,
	}
}

func TestCommentCreation(t *testing.T) {
	t.Parallel()

	t.Run("should create valid comment", func(t *testing.T) {
		suite := setupTest()

		suite.commentRepository.On("Save", mock.Anything, mock.Anything).Return(nil).Once()
		suite.userRepository.On("Find", mock.Anything, mock.Anything).
			Return(models.User{}, nil).
			Once()
		suite.postRepository.On("Find", mock.Anything, mock.Anything).
			Return(models.Post{}, nil).
			Once()

		err := suite.commentCreator.Run(context.Background(), "comment content", "1", "1")

		assert.Nil(t, err)
		suite.commentRepository.AssertExpectations(t)
		suite.userRepository.AssertExpectations(t)
		suite.postRepository.AssertExpectations(t)
	})

	t.Run("should fail when user does not exist", func(t *testing.T) {
		suite := setupTest()

		suite.userRepository.On("Find", mock.Anything, mock.Anything).
			Return(models.User{}, errors.ErrNotExistUser).Once()

		err := suite.commentCreator.Run(context.Background(), "comment content", "1", "1")

		assert.ErrorIs(t, err, errors.ErrNotExistUser)
		suite.userRepository.AssertExpectations(t)
	})

	t.Run("should fail when post does not exist", func(t *testing.T) {
		suite := setupTest()

		suite.userRepository.On("Find", mock.Anything, mock.Anything).
			Return(models.User{}, nil).
			Once()
		suite.postRepository.On("Find", mock.Anything, mock.Anything).
			Return(models.Post{}, errors.ErrNotExistPost).Once()

		err := suite.commentCreator.Run(context.Background(), "comment content", "1", "1")

		assert.ErrorIs(t, err, errors.ErrNotExistPost)
		suite.userRepository.AssertExpectations(t)
		suite.postRepository.AssertExpectations(t)
	})
}
