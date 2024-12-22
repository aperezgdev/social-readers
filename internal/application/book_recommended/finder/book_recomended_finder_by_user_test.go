package finder_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/aperezgdev/social-readers-api/internal/application/book_recommended/finder"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testSuite struct {
	mockBookRecommendedRepository *repository.MockBookRecommendedRepository
	bookRecommendedFinderByUser   *finder.BookRecommendedFinderByUser
}

func setupSuite() *testSuite {
	mockBookRecommendedRepository := &repository.MockBookRecommendedRepository{}
	bookRecommendedFinderByUser := finder.NewBookRecommendedFinderByUser(slog.Default(), mockBookRecommendedRepository)

	return &testSuite{
		mockBookRecommendedRepository,
		&bookRecommendedFinderByUser,
	}
}

func TestBookRecommendedSelection(t *testing.T) {
	t.Parallel()

	t.Run("should find recommended book", func(t *testing.T) {
		suite := setupSuite()

		suite.mockBookRecommendedRepository.On("FindByUser", mock.Anything, mock.Anything).
			Return([]models.BookRecommended{
				{
					Isbn:        "123123",
					Title:       "Hobbit",
					Description: "Description",
					Picture:     "picture",
					UserId:      "1",
				},
			}, nil)

		bookRecommendeds, err := suite.bookRecommendedFinderByUser.Run(context.Background(), "1")

		assert.Nil(t, err)
		assert.Equal(t, 1, len(bookRecommendeds))
		suite.mockBookRecommendedRepository.AssertExpectations(t)
	})
}