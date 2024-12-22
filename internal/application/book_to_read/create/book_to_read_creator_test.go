package create_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/social-readers-api/internal/application/book_to_read/create"
	"github.com/aperezgdev/social-readers-api/internal/domain/errors"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
)

type testSuite struct {
	mockBookToReadRepository *repository.MockBookToReadRepository
	mockUserRepository       *repository.MockUserRepository
	bookToReadCreator        *create.BookToReadCreator
}

func setupTest() *testSuite {
	mockBookToReadRepository := &repository.MockBookToReadRepository{}
	mockUserRepository := &repository.MockUserRepository{}
	bookToReadCreator := create.NewBookToReadCreator(
		slog.Default(),
		mockBookToReadRepository,
		mockUserRepository,
	)

	return &testSuite{
		mockBookToReadRepository,
		mockUserRepository,
		&bookToReadCreator,
	}
}

func TestBookToReadCreation(t *testing.T) {
	t.Parallel()

	t.Run("should create valid book to read", func(t *testing.T) {
		suite := setupTest()

		suite.mockUserRepository.On("Find", mock.Anything, mock.Anything).
			Return(models.User{}, nil).
			Once()
		suite.mockBookToReadRepository.On("Save", mock.Anything, mock.Anything).Return(nil).Once()

		err := suite.bookToReadCreator.Run(
			context.Background(),
			"123123",
			"Hobbit",
			"Description",
			"1",
			"picture",
		)

		assert.Nil(t, err)
		suite.mockUserRepository.AssertExpectations(t)
		suite.mockBookToReadRepository.AssertExpectations(t)
	})

	t.Run("should fail when user not exist", func(t *testing.T) {
		suite := setupTest()

		suite.mockUserRepository.On("Find", mock.Anything, mock.Anything).
			Return(models.User{}, errors.ErrNotExistUser)

		err := suite.bookToReadCreator.Run(
			context.Background(),
			"123123",
			"Hobbit",
			"Description",
			"2",
			"picture",
		)

		assert.ErrorIs(t, err, errors.ErrNotExistUser)
		suite.mockUserRepository.AssertExpectations(t)
	})
}
