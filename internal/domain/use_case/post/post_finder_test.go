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
	finder "github.com/aperezgdev/social-readers-api/internal/domain/use_case/post"
)

var (
	mockPostRepository = repository.MockPostRepository{}
	postFinder         = finder.NewPostFinder(slog.Default(), &mockPostRepository)
)

func Test_exist_post(t *testing.T) {
	postExpected, _ := models.NewPost("comment", "1")
	mockPostRepository.On("Find", mock.Anything, mock.Anything).Once().Return(postExpected, nil)

	post, err := postFinder.Run(context.Background(), string(postExpected.Id))
	assert.Nil(t, err)
	assert.Equal(t, postExpected, post)
}

func Test_non_exist_post(t *testing.T) {
	postExpected, _ := models.NewPost("comment", "1")
	mockPostRepository.On("Find", mock.Anything, mock.Anything).
		Once().
		Return(postExpected, errors.ErrNotExistPost)

	_, err := postFinder.Run(context.Background(), string(postExpected.Id))
	assert.ErrorIs(t, err, errors.ErrNotExistPost)
}
