package create_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/google/uuid"
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

		uuid, errUuid := uuid.NewV7()
		if errUuid != nil {
			t.Fatal(errUuid)
		}

		err := suite.bookRecommendedCreator.Run(
			context.Background(),
			"978-6-6795-0881-8",
			"Hobbit",
			"description",
			"picture",
			uuid.String(),
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

		uuid, errUuid := uuid.NewV7()
		if errUuid != nil {
			t.Fatal(errUuid)
		}
		err := suite.bookRecommendedCreator.Run(
			context.Background(),
			"0-4618-7203-X",
			"Hobbit",
			"description",
			"picture",
			uuid.String(),
		)

		assert.ErrorIs(t, err, errors.ErrNotExistUser)
		suite.mockUserRepository.AssertExpectations(t)
	})

	t.Run("should return validation error cause invalid book recommended title", func(t *testing.T) {
		suite := setupTest()

		uuid, errUuid := uuid.NewV7()
		if errUuid != nil {
			t.Fatal(errUuid)
		}
		err := suite.bookRecommendedCreator.Run(
			context.Background(),
			"978-6-6795-0881-8",
			" ",
			"description",
			"picture",
			uuid.String(),
		)

		assert.NotNil(t, err)
		validationError, ok := err.(errors.ValidationError)
		assert.True(t, ok)
		assert.Equal(t, "Title", validationError.Field)
	})
}
