package create_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/social-readers-api/internal/application/user/create"
	"github.com/aperezgdev/social-readers-api/internal/domain/errors"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
)

type testSuite struct {
	userRepository *repository.MockUserRepository
	userCreator    *create.UserCreator
}

func setupTest() *testSuite {
	mockUserRepo := &repository.MockUserRepository{}
	creator := create.NewUserCreator(
		slog.Default(),
		mockUserRepo,
	)

	return &testSuite{
		userRepository: mockUserRepo,
		userCreator:    &creator,
	}
}

func TestUserCreation(t *testing.T) {
	t.Run("should create valid user", func(t *testing.T) {
		suite := setupTest()
		suite.userRepository.On("Save", mock.Anything, mock.Anything).Return(nil).Once()

		err := suite.userCreator.Run(context.Background(), "John Doe", "john@doe.com", "picture")

		assert.Nil(t, err)
		suite.userRepository.AssertExpectations(t)
	})

	t.Run("should return validation error cause invalidad user name", func(t *testing.T) {
		suite := setupTest()
		err := suite.userCreator.Run(context.Background(), " ", "john@doe.com", "picture")

		assert.NotNil(t, err)
		validationError, ok := err.(errors.ValidationError)
		assert.True(t, ok)
		assert.Equal(t, "Name", validationError.Field)
	})
}
