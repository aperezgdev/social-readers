package create_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/social-readers-api/internal/application/book_recommended/create"
	"github.com/aperezgdev/social-readers-api/internal/domain/errors"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
)

type testSuite struct {
	mockBookRecommendedRepository *repository.MockBookRecommendedRepository
	mockUserRepository            *repository.MockUserRepository
	bookRecommendedCreator        *create.BookRecommendedCreator
}

func setupTest() *testSuite {
	mockBookRecommendedRepository := &repository.MockBookRecommendedRepository{}
	mockUserRepository := &repository.MockUserRepository{}
	bookRecommendedCreator := create.NewBookRecommendedCreator(
		slog.Default(),
		mockBookRecommendedRepository,
		mockUserRepository,
	)

	return &testSuite{
		mockBookRecommendedRepository,
		mockUserRepository,
		&bookRecommendedCreator,
	}
}

func TestBookRecommendedCreation(t *testing.T) {
	t.Parallel()

	t.Run("should create valid book recommended", func(t *testing.T) {
		suite := setupTest()

		suite.mockUserRepository.On("Find", mock.Anything, mock.Anything).
			Return(models.User{}, nil).
			Once()
		suite.mockBookRecommendedRepository.On("Save", mock.Anything, mock.Anything).
			Return(nil).
			Once()

		err := suite.bookRecommendedCreator.Run(
			context.Background(),
			"123123",
			"Hobbit",
			"description",
			"picture",
			"1",
		)

		assert.Nil(t, err)
		suite.mockUserRepository.AssertExpectations(t)
		suite.mockBookRecommendedRepository.AssertExpectations(t)
	})

	t.Run("should fail when user does not exist", func(t *testing.T) {
		suite := setupTest()

		suite.mockUserRepository.On("Find", mock.Anything, mock.Anything).
			Return(models.User{}, errors.ErrNotExistUser).
			Once()

		err := suite.bookRecommendedCreator.Run(
			context.Background(),
			"123123",
			"Hobbit",
			"description",
			"picture",
			"1",
		)

		assert.ErrorIs(t, err, errors.ErrNotExistUser)
		suite.mockUserRepository.AssertExpectations(t)
	})
}
