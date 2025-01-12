package create_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/social-readers-api/internal/application/post/create"
	"github.com/aperezgdev/social-readers-api/internal/domain/errors"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
)

var (
	mockPostRepository = repository.MockPostRepository{}
	mockUserRepository = repository.MockUserRepository{}
	postCreator        = create.NewPostCreator(
		slog.Default(),
		&mockPostRepository,
		&mockUserRepository,
	)
)

func Test_create_valid_post(t *testing.T) {
	t.Parallel()

	t.Run("should create a valid post without error", func(t *testing.T) {
		uuid, errUuid := uuid.NewV7()
		if errUuid != nil {
			t.Error("Fail trying to create a uuid for testing")
		}
		mockPostRepository.On("Save", mock.Anything, mock.Anything).Once().Return(nil)
		mockUserRepository.On("Find", mock.Anything, mock.Anything).Once().Return(models.User{}, nil)
		err := postCreator.Run(context.Background(), "Amazing book!", "978-6-6795-0881-8", uuid.String())

		assert.Nil(t, err)
	})

	t.Run("should return not exist user", func(t *testing.T) {
		uuid, errUuid := uuid.NewV7()
		if errUuid != nil {
			t.Error("Fail trying to create a uuid for testing")
		}
		mockUserRepository.On("Find", mock.Anything, mock.Anything).
		Once().
		Return(models.User{}, errors.ErrNotExistUser)
		err := postCreator.Run(context.Background(), "Amazing book!", "978-6-6795-0881-8", uuid.String())

		assert.ErrorIs(t, err, errors.ErrNotExistUser)
	})

	t.Run("should return a validation error on invalid post", func(t *testing.T) {
		err := postCreator.Run(context.Background(), "Amazing book!", "978-6-6795-0881-8","1")
		validationError, ok := err.(errors.ValidationError)
		assert.True(t, ok)
		assert.Equal(t, "postedBy", validationError.Field)
	})
}
