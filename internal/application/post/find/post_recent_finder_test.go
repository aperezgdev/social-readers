package finder_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	finder "github.com/aperezgdev/social-readers-api/internal/application/post/find"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
)

type testSuite struct {
	mockPostRepository *repository.MockPostRepository
	postRecentFinder   finder.PostRecentFinder
}

func setupSuite() *testSuite {
	mockPostRepository := &repository.MockPostRepository{}
	postRecentFinder := finder.NewPostRecentFinder(slog.Default(), mockPostRepository)

	return &testSuite{
		mockPostRepository,
		*postRecentFinder,
	}
}

func TestPostRecentSelection(t *testing.T) {
	t.Parallel()

	t.Run("should find recent post", func(t *testing.T) {
		suite := setupSuite()

		suite.mockPostRepository.On("FindRecent", mock.Anything, mock.Anything).
			Return(make([]models.Post, 10), nil)

		posts, err := suite.postRecentFinder.Run(context.Background())

		assert.Nil(t, err)
		assert.True(t, len(posts) == 10)
		suite.mockPostRepository.AssertExpectations(t)
	})
}
