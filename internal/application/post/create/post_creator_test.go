package create_test

import (
	"context"
	"log/slog"
	"testing"

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
	mockPostRepository.On("Save", mock.Anything, mock.Anything).Once().Return(nil)
	mockUserRepository.On("Find", mock.Anything, mock.Anything).Once().Return(models.User{}, nil)
	err := postCreator.Run(context.Background(), "Amazing book!", "1")

	assert.Nil(t, err)
}

func Test_create_post_non_existing_user(t *testing.T) {
	mockUserRepository.On("Find", mock.Anything, mock.Anything).
		Once().
		Return(models.User{}, errors.ErrNotExistUser)
	err := postCreator.Run(context.Background(), "Amazing book!", "2")

	assert.ErrorIs(t, err, errors.ErrNotExistUser)
}

func Test_create_an_invalid_post(t *testing.T) {
}
