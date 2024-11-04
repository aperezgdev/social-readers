package create_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/social-readers-api/internal/application/user/create"
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

	// You can add more test cases here as needed, for example:
	/*
		t.Run("should fail with invalid email", func(t *testing.T) {
			suite := setupTest()

			err := suite.userCreator.Run(context.Background(), "John Doe", "invalid-email", "picture")

			assert.Error(t, err)
			suite.userRepository.AssertNotCalled(t, "Save")
		})
	*/
}
