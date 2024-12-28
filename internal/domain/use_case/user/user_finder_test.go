package finder_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/social-readers-api/internal/domain/errors"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	finder "github.com/aperezgdev/social-readers-api/internal/domain/use_case/user"
)

var (
	mockUserRepository = repository.MockUserRepository{}
	userFinder         = finder.NewUserFinder(slog.Default(), &mockUserRepository)
)

func Test_exist_user(t *testing.T) {
	userExpected, _ := models.NewUser("John", "picture", "john@doe.com")
	mockUserRepository.On("Find", mock.Anything, mock.Anything).
		Once().
		Return(userExpected, nil)

	user, err := userFinder.Run(context.Background(), string(userExpected.Id))

	assert.Nil(t, err)
	assert.Equal(t, user, userExpected)
}

func Test_non_exist_user(t *testing.T) {
	userExpected, _ := models.NewUser("John", "picture", "john@doe.com")
	mockUserRepository.On("Find", mock.Anything, mock.Anything).
		Once().
		Return(userExpected, errors.ErrNotExistUser)

	_, err := userFinder.Run(context.Background(), string(userExpected.Id))

	assert.ErrorIs(t, err, errors.ErrNotExistUser)
}
